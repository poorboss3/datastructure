package leetcode

import (
	"reflect"
	"testing"
)

func Test_alertName(t *testing.T) {
	type args struct {
		keyName []string
		keyTime []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				keyName: []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
				keyTime: []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
			},
			want: []string{"daniel"},
		},
		{
			name: "test 2",
			args: args{
				keyName: []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
				keyTime: []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
			},
			want: []string{"bob"},
		},
		{
			name: "test 3",
			args: args{
				keyName: []string{"john", "john", "john"},
				keyTime: []string{"23:58", "23:59", "00:01"},
			},
			want: []string{},
		},
		{
			name: "test 4",
			args: args{
				keyName: []string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
				keyTime: []string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
			},
			want: []string{"clare", "leslie"},
		},
		{
			name: "test 4",
			args: args{
				keyName: []string{"b", "b", "b", "b", "b", "b"},
				keyTime: []string{"00:52", "10:59", "17:16", "00:36", "01:26", "22:42"},
			},
			want: []string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alertName(tt.args.keyName, tt.args.keyTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alertName() = %v, want %v", got, tt.want)
			}
		})
	}
}
