package datacenter

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sanguohot/medichain/etc"
	"github.com/sanguohot/medichain/util"
	"github.com/sanguohot/medichain/zap"
	"io/ioutil"
	"os"
)
type FileAddLog struct {
	FileUuid	string
	OwnerUuid	string
	UploaderUuid	string
	OrgUuid	string
	FileTypeHash	string
	FileType	string
	FileDesc	string
	Keccak256Hash	string
	Sha256Hash	string
	CreateTime	uint64
	Signature	string
	Signer	string
	BlockNum	uint64
	BlockHash	string
	TransactionHash	string
	ContractAddress	string
}
func init()  {
	if !util.FilePathExist(etc.GetSqliteFilePath()) {
		_, err := os.Create(etc.GetSqliteFilePath())
		if err != nil {
			zap.Logger.Fatal(err.Error())
		}
	}
	// 确保sql可以重复执行 也就是包含IF NOT EXISTS
	if !util.FilePathExist(etc.GetSqliteFileAddLogPath()) {
		zap.Logger.Fatal(fmt.Errorf("sqlite: not found %s", etc.GetSqliteFileAddLogPath()).Error())
	}
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	defer db.Close()
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	data, err := ioutil.ReadFile(etc.GetSqliteFileAddLogPath())
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
	_, err = db.Exec(string(data))
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
}

func SqliteSetFileAddLogList(fl []FileAddLog) error {
	if fl == nil || len(fl) == 0 {
		return nil
	}
	zap.Sugar.Infof("sqlite: now insert []FileAddLog %d", len(fl))
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	defer db.Close()
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("INSERT  OR IGNORE INTO tbl_file_add_event_log(FileUuid, OwnerUuid, UploaderUuid, OrgUuid, FileTypeHash, FileType" +
		", FileDesc, Keccak256Hash, Sha256Hash, CreateTime, Signature, Signer, BlockNum, BlockHash, TransactionHash" +
			", ContractAddress) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	var res sql.Result
	for _, f := range fl {
		res, err = stmt.Exec(f.FileUuid, f.OwnerUuid, f.UploaderUuid, f.OrgUuid, f.FileTypeHash, f.FileType, f.FileDesc, f.Keccak256Hash, f.Sha256Hash, f.CreateTime,
			f.Signature, f.Signer, f.BlockNum, f.BlockHash, f.TransactionHash, f.ContractAddress)
		if err != nil {
			return err
		}
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func SqliteGetFileAddLogList(fileUuid, orgUuid, ownerUuid string, fromTime, toTime, start, limit uint64) (error, []FileAddLog) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	defer db.Close()
	if err != nil {
		return err, nil
	}
	sql := getFileAddLogQuerySql(false, fileUuid, orgUuid, ownerUuid, fromTime, toTime, start, limit)
	zap.Sugar.Infof("sqlite: %s", sql)
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return err, nil
	}
	var fl []FileAddLog
	for rows.Next() {
		f := FileAddLog{}
		err = rows.Scan(&f.FileUuid, &f.OwnerUuid, &f.UploaderUuid, &f.OrgUuid, &f.FileTypeHash, &f.FileType, &f.FileDesc, &f.Keccak256Hash, &f.Sha256Hash, &f.CreateTime,
			&f.Signature, &f.Signer, &f.BlockNum, &f.BlockHash, &f.TransactionHash, &f.ContractAddress)
		if err != nil {
			return err, nil
		}
		fl = append(fl, f)
	}
	return nil, fl
}

func SqliteGetFileAddLogMaxBlockNum() (error, uint64) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	defer db.Close()
	if err != nil {
		return err, 0
	}
	sql := "SELECT MAX(BlockNum) FROM tbl_file_add_event_log"
	//fmt.Printf("sqlite: %s\n", sql)
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return err, 0
	}
	// 需要设置为指针, 并且判断是否为nil
	var blockNum *uint64
	for rows.Next() {
		err = rows.Scan(&blockNum)
		if err != nil {
			return err, 0
		}
		if blockNum == nil {
			return nil, 0
		}
		break
	}
	return nil, *blockNum
}

func SqliteGetFileAddLogTotal(fileUuid, orgUuid, ownerUuid string, fromTime, toTime uint64) (error, uint64) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	defer db.Close()
	if err != nil {
		return err, 0
	}
	sql := getFileAddLogQuerySql(true, fileUuid, orgUuid, ownerUuid, fromTime, toTime, 0, 0)
	zap.Sugar.Infof("sqlite: %s", sql)
	rows, err := db.Query(sql)
	defer rows.Close()
	if err != nil {
		return err, 0
	}
	// 需要设置为指针, 并且判断是否为nil
	var count *uint64
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return err, 0
		}
		if count == nil {
			return nil, 0
		}
		break
	}
	return nil, *count
}

func getFileAddLogQuerySql(isCount bool, fileUuid, orgUuid, ownerUuid string, fromTime, toTime, start, limit uint64) string {
	sql := "SELECT * FROM tbl_file_add_event_log WHERE 1=1"
	if isCount  {
		sql = "SELECT COUNT(*) FROM tbl_file_add_event_log WHERE 1=1"
	}
	if ownerUuid != "" {
		sql = fmt.Sprintf("%s AND OwnerUuid=\"%s\"", sql, ownerUuid)
	}
	if orgUuid != "" {
		sql = fmt.Sprintf("%s AND OrgUuid=\"%s\"", sql, orgUuid)
	}
	if fileUuid != "" {
		sql = fmt.Sprintf("%s AND FileUuid=\"%s\"", sql, fileUuid)
	}
	if fromTime > 0 {
		sql = fmt.Sprintf("%s AND CreateTime>=%d", sql, fromTime)
	}
	if toTime > 0 {
		sql = fmt.Sprintf("%s AND CreateTime<=%d", sql, toTime)
	}
	if !isCount {
		sql = fmt.Sprintf("%s ORDER BY CreateTime DESC", sql)
	}
	if limit > 0 {
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, limit, start)
	}
	return sql
}