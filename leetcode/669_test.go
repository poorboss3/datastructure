package leetcode

import (
	"reflect"
	"testing"
)

var tree = initTree()

func initTree() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 1}
	return root
}

func Test_trimBST4(t *testing.T) {
	type args struct {
		root *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: tree,
				low:  1,
				high: 1,
			},
			want: &TreeNode{Val: 1, Left: nil, Right: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST1(tt.args.root, tt.args.low, tt.args.high); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST4() = %v, want %v", got, tt.want)
			}
		})
	}
}
