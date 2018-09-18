package service

import (
	"reflect"
	"testing"

	"github.com/sanguohot/medichain/datacenter"
)

func TestSetFileAddLogList(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetFileAddLogList(); (err != nil) != tt.wantErr {
				t.Errorf("SetFileAddLogList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileAddLogList(t *testing.T) {
	type args struct {
		idCartNo   string
		orgUuidStr string
		startStr   string
		limitStr   string
	}
	tests := []struct {
		name    string
		args    args
		want    *FileAddLogSimpleAction
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", "", ""}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetFileAddLogList(tt.args.idCartNo, tt.args.orgUuidStr, tt.args.startStr, tt.args.limitStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileAddLogList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileAddLogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileAddLogDetail(t *testing.T) {
	type args struct {
		fileUuidStr string
	}
	tests := []struct {
		name    string
		args    args
		want    *datacenter.FileAddLog
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := GetFileAddLogDetail(tt.args.fileUuidStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileAddLogDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileAddLogDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
