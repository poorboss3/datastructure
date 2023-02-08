package leetcode

import (
	"reflect"
	"testing"
)

func Test_ambiguousCoordinates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{s: "(123)"},
			want: []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"},
		},
		{
			name: "test2",
			args: args{s: "(00011)"},
			want: []string{"(0.001, 1)", "(0, 0.011)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ambiguousCoordinates(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ambiguousCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
