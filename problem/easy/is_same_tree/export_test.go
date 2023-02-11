package is_same_tree

import (
	"testing"

	"github.com/supermarine1377/leetcode/types"
)

type args struct {
	p *TreeNode
	q *TreeNode
}

type testcase struct {
	name string
	args args
	want bool
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				p: nil,
				q: nil,
			},
			want: true,
		},
		{
			name: "2nd",
			args: args{
				p: &types.TreeNode{
					Val: 1,
					Left: &types.TreeNode{
						Val: 2,
					},
					Right: nil,
				},
				q: &types.TreeNode{
					Val: 1,
					Left: &types.TreeNode{
						Val: 2,
					},
					Right: nil,
				},
			},
			want: true,
		},
		{
			name: "3rd",
			args: args{
				p: &types.TreeNode{
					Val: 1,
					Left: &types.TreeNode{
						Val: 2,
					},
					Right: nil,
				},
				q: &types.TreeNode{
					Val: 1,
					Left: &types.TreeNode{
						Val: 1,
					},
					Right: nil,
				},
			},
			want: false,
		},
	}
}

func Test_isSameTree(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSameTreeRecursively(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTreeRecursively(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
