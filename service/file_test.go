package service

import (
	"reflect"
	"testing"
)

func TestAddFile(t *testing.T) {
	type args struct {
		ownerUuidStr  string
		orgUuidStr    string
		addressStr    string
		password      string
		fileType      string
		fileDesc      string
		file          []byte
		sha256HashStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", "", "", "", "", nil, ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := AddFile(tt.args.ownerUuidStr, tt.args.orgUuidStr, tt.args.addressStr, tt.args.password, tt.args.fileType, tt.args.fileDesc, tt.args.file, tt.args.sha256HashStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddFileSign(t *testing.T) {
	type args struct {
		fileUuidStr      string
		addressStr       string
		password         string
		keccak256HashStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileAddSignAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", "", ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := AddFileSign(tt.args.fileUuidStr, tt.args.addressStr, tt.args.password, tt.args.keccak256HashStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddFileSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFileSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFile(t *testing.T) {
	type args struct {
		fileUuidStr string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetFile(tt.args.fileUuidStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileSignerAndDataList(t *testing.T) {
	type args struct {
		fileUuidStr string
		startStr    string
		limitStr    string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileSignerAndDataAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetFileSignerAndDataList(tt.args.fileUuidStr, tt.args.startStr, tt.args.limitStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileSignerAndDataList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileSignerAndDataList() = %v, want %v", got, tt.want)
			}
		})
	}
}
