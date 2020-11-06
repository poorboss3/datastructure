package array

import (
	"testing"
)

func TestMaxSubsequenceSum(t *testing.T) {
	type args struct {
		a []int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{[]int{-4, 3, 5, -2, -1, 2, -6, -2}, 7},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubsequenceSum(tt.args.a, tt.args.N); got != tt.want {
				t.Errorf("array is %v,len is %v,MaxSubsequenceSum() = %v, want %v", tt.args.a, tt.args.N, got, tt.want)
			}
		})
	}
}

func TestMaxSubsequenceSumDi(t *testing.T) {
	type args struct {
		a []int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test2",
			args: args{[]int{-4, 3, 5, -2, -1, 2, -6, -2}, 7},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubsequenceSumDi(tt.args.a, tt.args.N); got != tt.want {
				t.Errorf("MaxSubsequenceSumDi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubsequenceSumDia(t *testing.T) {
	type args struct {
		a []int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test3",
			args: args{[]int{-4, 3, 5, -2, -1, 2, -6, -2}, 7},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubsequenceSumDia(tt.args.a, tt.args.N); got != tt.want {
				t.Errorf("MaxSubsequenceSumDia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxSubsequenceSumOn(t *testing.T) {
	type args struct {
		a []int
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test3",
			args: args{[]int{-4, 3, 5, -2, -1, 2, -6, -2}, 7},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubsequenceSumOn(tt.args.a, tt.args.N); got != tt.want {
				t.Errorf("MaxSubsequenceSumOn() = %v, want %v", got, tt.want)
			}
		})
	}
}
