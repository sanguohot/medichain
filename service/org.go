package service

import (
	"github.com/ethereum/go-ethereum/common"
	"crypto"
	"github.com/google/uuid"
)

type Org struct {
	Uuid uuid.UUID
	Address common.Address
	PublicKey crypto.PublicKey
	nameHash common.Hash
}

func AddOrg(orgUuid uuid.UUID, idCartNo string) (error, *Org, *common.Hash) {
	return nil, nil, nil
}

func GetOrgByUuid(uuid uuid.UUID) (error, *Org) {
	return nil, nil
}

func GetOrgByName(name string) (error, *Org) {
	return nil, nil
}

func GetOrgByAddress(address common.Address) (error, *Org) {
	return nil, nil
}