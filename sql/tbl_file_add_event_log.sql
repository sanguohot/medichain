CREATE TABLE IF NOT EXISTS `tbl_file_add_event_log` (
	`FileUuid`	TEXT NOT NULL,
	`OwnerUuid`	TEXT NOT NULL,
	`UploaderUuid`	TEXT NOT NULL,
	`OrgUuid`	TEXT,
	`FileTypeHash`	TEXT NOT NULL,
	`FileType`	TEXT NOT NULL,
	`FileDesc`	TEXT,
	`Keccak256Hash`	TEXT NOT NULL UNIQUE,
	`Sha256Hash`	TEXT NOT NULL UNIQUE,
	`CreateTime`	INTEGER NOT NULL,
	`Signature`	TEXT NOT NULL,
	`Signer`	TEXT NOT NULL,
	`BlockNum`	INTEGER NOT NULL,
	`BlockHash`	TEXT NOT NULL,
	`TransactionHash`	TEXT NOT NULL UNIQUE,
	`ContractAddress`	TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS `idx1` ON `tbl_file_add_event_log` (`OwnerUuid`);
CREATE INDEX IF NOT EXISTS `idx2` ON `tbl_file_add_event_log` (`FileTypeHash`);
CREATE INDEX IF NOT EXISTS `idx3` ON `tbl_file_add_event_log` (`UploaderUuid`);
CREATE INDEX IF NOT EXISTS `idx4` ON `tbl_file_add_event_log` (`ContractAddress`);
CREATE INDEX IF NOT EXISTS `idx5` ON `tbl_file_add_event_log` (`OrgUuid`);