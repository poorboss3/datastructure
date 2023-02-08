package leetcode

import (
	"reflect"
	"testing"
)

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				text: "  this   is  a sentence ",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitSpace(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				text: "  this   is  a sentence ",
			},
			want:  []string{"this", "is", "a", "sentence"},
			want1: 9,
		},
		{
			name: "",
			args: args{
				text: " practice   makes   perfect",
			},
			want:  []string{"practice", "makes", "perfect"},
			want1: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitSpace(tt.args.text)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitSpace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitSpace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
