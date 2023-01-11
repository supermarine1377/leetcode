package maximum_depth_of_binary_tree

import "testing"

func Test_maxDepthRecursively(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1st",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: 2,
		},
		{
			name: "2nd",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: nil,
					},
					Right: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 5,
							},
							Right: nil,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			}, want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthRecursively(tt.args.root); got != tt.want {
				t.Errorf("maxDepthRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
