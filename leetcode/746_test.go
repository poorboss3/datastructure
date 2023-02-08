package leetcode

import "testing"

func Test_orderOfLargesPlusSign(t *testing.T) {
	type args struct {
		n     int
		mines [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				n:     5,
				mines: [][]int{{1, 0}, {1, 4}, {2, 4}, {3, 2}, {4, 0}, {4, 3}},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				n:     10,
				mines: [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 4}, {0, 5}, {0, 8}, {0, 9}, {1, 0}, {1, 1}, {1, 3}, {1, 5}, {1, 6}, {1, 7}, {1, 9}, {2, 0}, {2, 1}, {2, 2}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}, {3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5}, {3, 7}, {3, 8}, {4, 0}, {4, 1}, {4, 2}, {4, 4}, {4, 8}, {4, 9}, {5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 6}, {5, 8}, {5, 9}, {6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 5}, {6, 7}, {6, 8}, {6, 9}, {7, 0}, {7, 1}, {7, 6}, {7, 7}, {7, 8}, {7, 9}, {8, 0}, {8, 1}, {8, 2}, {8, 5}, {8, 9}, {9, 1}, {9, 3}, {9, 4}, {9, 5}, {9, 6}, {9, 8}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := orderOfLargestPlusSign(tt.args.n, tt.args.mines); got != tt.want {
				t.Errorf("orderOfLargesPlusSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
