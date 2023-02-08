package leetcode

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//TODO: Add test cases.
		{
			name: "98368",
			args: args{
				num: 98368,
			},
			want: 98863,
		},
		{
			name: "9933",
			args: args{
				num: 2736,
			},
			want: 7236,
		},
		{
			name: "9933",
			args: args{
				num: 9933,
			},
			want: 9933,
		},
		{
			name: "21",
			args: args{
				num: 12,
			},
			want: 21,
		},
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: 1,
		},
		{
			name: "10",
			args: args{
				num: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
