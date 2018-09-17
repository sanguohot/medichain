package datacenter

import (
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqliteSetFileAddLogList(t *testing.T) {
	type args struct {
		fl []FileAddLog
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: args{fl: nil}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SqliteSetFileAddLogList(tt.args.fl); (err != nil) != tt.wantErr {
				t.Errorf("SqliteSetFileAddLogList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSqliteGetFileAddLogList(t *testing.T) {
	type args struct {
		fileUuid  string
		orgUuid   string
		ownerUuid string
		start     uint64
		limit     uint64
	}
	tests := []struct {
		name    string
		args    args
		want    []FileAddLog
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", args: args{fileUuid: "", orgUuid: "", ownerUuid: "", start: 0, limit: 10}, want: nil, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SqliteGetFileAddLogList(tt.args.fileUuid, tt.args.orgUuid, tt.args.ownerUuid, tt.args.start, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteGetFileAddLogList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqliteGetFileAddLogList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqliteGetFileAddLogMaxBlockNum(t *testing.T) {
	tests := []struct {
		name    string
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "", want: 0, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := SqliteGetFileAddLogMaxBlockNum()
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteGetFileAddLogMaxBlockNum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == 0 {
				return
			}
			if got != tt.want {
				t.Errorf("SqliteGetFileAddLogMaxBlockNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSqliteGetFileAddLogTotal(t *testing.T) {
	type args struct {
		fileUuid  string
		orgUuid   string
		ownerUuid string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{"", "", ""}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, got := SqliteGetFileAddLogTotal(tt.args.fileUuid, tt.args.orgUuid, tt.args.ownerUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("SqliteGetFileAddLogTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == 0 {
				return
			}
			if got != tt.want {
				t.Errorf("SqliteGetFileAddLogTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
