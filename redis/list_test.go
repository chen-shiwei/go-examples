package redis

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestLPush(t *testing.T) {
	type args struct {
		key string
		val []interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantRows int64
		wantErr  bool
	}{
		{name: "list LPush", args: args{
			key: "1",
			val: []interface{}{[]byte("a")},
		}, wantRows: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, err := LPush(tt.args.key, tt.args.val...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LPush() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRows != tt.wantRows {
				t.Errorf("LPush() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
		})
	}
}

func TestLPushX(t *testing.T) {
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name     string
		args     args
		wantRows int64
		wantErr  bool
	}{
		{name: "key1", args: args{
			key: "1",
			val: "a",
		}, wantRows: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRows, err := LPushX(tt.args.key, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("LPushX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRows != tt.wantRows {
				t.Errorf("LPushX() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{
			numbers: []int{5, 25, 75},
			target:  100,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPVowels(t *testing.T) {
	s := "aeiouAEIOU"
	l := len(s)
	u := []uint8{97, 101, 105, 111, 117, 65, 69, 73, 79, 85}
	for i := 0; i < l-1; i++ {
		if s[i] == u[i] {
			t.Logf("i:%v v:%v\n", i, s[i])
		}
	}
}

//97, 101, 105, 111, 117, 65, 69, 73, 79, 85

func TestReverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "leetcode"}, want: "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseVowels(tt.args.s); got != tt.want {
				t.Errorf("ReverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
