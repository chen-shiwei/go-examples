package redis

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	t.Run("TestStringSet", func(t *testing.T) {
		TestStringSet(t)
	})
	t.Run("TestStringGet", func(t *testing.T) {
		TestStringGet(t)
	})
	t.Run("TestStringDel", func(t *testing.T) {
		TestStringDel(t)
	})
}

func TestStringSet(t *testing.T) {
	type args struct {
		key        string
		val        interface{}
		expiration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "string set", args: args{
				key:        "key1",
				val:        "value1",
				expiration: time.Minute * 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StringSet(tt.args.key, tt.args.val, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("StringSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringGet(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantVal string
		wantErr bool
	}{
		{name: "string get", args: args{key: "key1"}, wantVal: "value1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, err := StringGet(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotVal != tt.wantVal {
				t.Errorf("StringGet() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestStringDel(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		args     args
		wantRows int64
		wantErr  bool
	}{
		{name: "key1", args: args{key: "key1"}, wantRows: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, err := StringDel(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringDel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRows != tt.wantRows {
				t.Errorf("StringDel() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
		})
	}
}
