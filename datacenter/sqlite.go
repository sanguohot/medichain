package datacenter

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"medichain/etc"
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
func SqliteSetFileAddLogList(fl []FileAddLog) error {
	if len(fl) == 0 {
		return nil
	}
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO tbl_file_add_event_log(FileUuid, OwnerUuid, UploaderUuid, OrgUuid, FileTypeHash, FileType" +
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

func SqliteGetFileAddLogList(fileUuid, ownerUuid string, start, limit uint64) (error, []FileAddLog) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	if err != nil {
		return err, nil
	}
	defer db.Close()
	var (
		ownerQuery string
		fileQuery string
		limitQuery string
	)
	if ownerUuid != "" {
		ownerQuery = fmt.Sprintf("AND OwnerUuid=%s", ownerUuid)
	}
	if fileUuid != "" {
		fileQuery = fmt.Sprintf("AND FileUuid=\"%s\"", fileUuid)
	}
	if limit > 0 {
		limitQuery = fmt.Sprintf("LIMIT %d OFFSET %d", limit, start)
	}
	sql := fmt.Sprintf("SELECT * FROM tbl_file_add_event_log WHERE 1=1 %s %s ORDER BY CreateTime DESC %s", ownerQuery, fileQuery, limitQuery)
	fmt.Println(sql)
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

func SqliteGetMaxBlockNum() (error, uint64) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	if err != nil {
		return err, 0
	}
	defer db.Close()
	sql := "SELECT MAX(BlockNum) FROM tbl_file_add_event_log"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	defer rows.Close()
	fmt.Println(rows.Columns())
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