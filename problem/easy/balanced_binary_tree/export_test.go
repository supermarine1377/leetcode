package balanced_binary_tree

import (
	"testing"

	"github.com/supermarine1377/leetcode/types"
)

type args struct {
	root *TreeNode
}

type testcase struct {
	name string
	args args
	want bool
}

func getTestCase() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				root: &TreeNode{
					Left: &TreeNode{},
					Right: &TreeNode{
						Left:  &TreeNode{},
						Right: &TreeNode{},
					},
				},
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				root: &TreeNode{
					Left: &TreeNode{},
					Right: &TreeNode{
						Left: &TreeNode{},
						Right: &TreeNode{
							Left: &types.TreeNode{
								Left: &TreeNode{},
							},
						},
					},
				},
			},
			want: false,
		},
	}
}

func Test_isBalanced(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBalanced_ON(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced_ON(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
