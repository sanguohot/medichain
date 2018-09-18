package service

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddOrg(t *testing.T) {
	type args struct {
		name     string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *OrgAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "123456"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := AddOrg(tt.args.name, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddOrg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddOrg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrgByNameHash(t *testing.T) {
	type args struct {
		hash common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    *OrgAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{common.Hash{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetOrgByNameHash(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrgByNameHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrgByNameHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
