package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "sort1", args: args{array: []int{7, 1, 4, 2, 5, 3, 6}}, want: []int{1, 2, 3, 4, 5, 6, 7}},
		{name: "sort2", args: args{array: []int{7, 1, 4, 8, 9, 2, 5, 3, 6}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
