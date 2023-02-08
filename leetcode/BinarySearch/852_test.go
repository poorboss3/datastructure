package binarysearch

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{arr: []int{3, 5, 3, 2, 0}},
			want: 1,
		},
		{
			name: "test2",
			args: args{arr: []int{0, 2, 1, 0}},
			want: 1,
		},
		{
			name: "test3",
			args: args{arr: []int{0, 10, 5, 2}},
			want: 1,
		},
		{
			name: "test4",
			args: args{arr: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
