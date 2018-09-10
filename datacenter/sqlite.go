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
	stmt, err := db.Prepare("INSERT INTO tbl_file_add_event_log(FileUuid, OwnerUuid, UploaderUuid, OrgUuid, FileType" +
		", FileDesc, Keccak256Hash, Sha256Hash, CreateTime, Signature, Signer, BlockNum, BlockHash, TransactionHash" +
			", ContractAddress) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	var res sql.Result
	for _, f := range fl {
		res, err = stmt.Exec(f.FileUuid, f.OwnerUuid, f.UploaderUuid, f.OrgUuid, f.FileType, f.FileDesc, f.Keccak256Hash, f.Sha256Hash, f.CreateTime,
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

func SqliteGetFileAddLogList(ownerUuid string) (error, []FileAddLog) {
	db, err := sql.Open("sqlite3", etc.GetSqliteFilePath())
	if err != nil {
		return err, nil
	}
	defer db.Close()
	ownerQuery := ""
	if ownerUuid != "" {
		ownerQuery = fmt.Sprintf(" and OwnerUuid=%s", ownerUuid)
	}
	sql := fmt.Sprintf(`
	SELECT * FROM tbl_file_add_event_log
	WHERE 1=1
	%s
	ORDER BY CreateTime DESC
`, ownerQuery)
	rows, err := db.Query(sql)
	if err != nil {
		return err, nil
	}
	defer rows.Close()

	var fl []FileAddLog
	for rows.Next() {
		f := FileAddLog{}
		err = rows.Scan(&f.FileUuid, &f.OwnerUuid, &f.UploaderUuid, &f.OrgUuid, &f.FileType, &f.FileDesc, &f.Keccak256Hash, &f.Sha256Hash, &f.CreateTime,
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
	rows, err := db.Query(sql)
	if err != nil {
		return err, 0
	}
	defer rows.Close()
	var blockNum uint64
	for rows.Next() {
		err = rows.Scan(&blockNum)
		if err != nil {
			return err, 0
		}
		break
	}
	return nil, blockNum
}