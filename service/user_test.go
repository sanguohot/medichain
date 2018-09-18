package service

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddUser(t *testing.T) {
	type args struct {
		orgUuidStr string
		idCartNo   string
		password   string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", "123456"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := AddUser(tt.args.orgUuidStr, tt.args.idCartNo, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUserByIdCartNoHash(t *testing.T) {
	type args struct {
		hash common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    *UserAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{common.Hash{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetUserByIdCartNoHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByIdCartNoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByIdCartNoHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
