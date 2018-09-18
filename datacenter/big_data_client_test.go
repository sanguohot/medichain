package datacenter

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestUploadToBigDataCenter(t *testing.T) {
	type args struct {
		fileBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UploadToBigDataCenter(tt.args.fileBytes); (err != nil) != tt.wantErr {
				t.Errorf("UploadToBigDataCenter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDownloadFromBigDataCenter(t *testing.T) {
	type args struct {
		sha256Hash common.Hash
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{common.Hash{}}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadFromBigDataCenter(tt.args.sha256Hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFromBigDataCenter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadFromBigDataCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFolderInBigDataCenter(t *testing.T) {
	type args struct {
		folderName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{""}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFolderInBigDataCenter(tt.args.folderName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFolderInBigDataCenter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateFolderInBigDataCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFilesInFolder(t *testing.T) {
	type args struct {
		folderCode string
	}
	tests := []struct {
		name    string
		args    args
		want    []BigDataItem
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{""}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFilesInFolder(tt.args.folderCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilesInFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilesInFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
