package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type File struct {
	Uuid uuid.UUID
	Keccak256Hash common.Hash
	Sha256Hash common.Hash
}

func AddFile(ownerUuid, uploaderUuid uuid.UUID, keystorePwd, fileType, fileDesc string) (error, *File, *common.Hash) {
	return nil, nil, nil
}

func GetFileByUuid(uuid uuid.UUID) (error, *File) {
	return nil, nil
}

func GetFileByKeccak256Hash(keccak256Hash common.Hash) (error, *File) {
	return nil, nil
}

func GetFileSha256Hash(sha256Hash common.Hash) (error, *File) {
	return nil, nil
}