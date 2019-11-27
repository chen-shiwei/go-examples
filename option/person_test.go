package option

import (
	"reflect"
	"testing"
)

func TestNewPerson(t *testing.T) {
	type args struct {
		ops []Option
	}
	tests := []struct {
		name string
		args args
		want *Person
	}{
		{
			name: "alice,16",
			args: args{
				ops: []Option{
					WithAge(16),
					WithName("alice"),
				},
			},
			want: &Person{options{
				Name: "alice",
				Age:  16,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPerson(tt.args.ops...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}
