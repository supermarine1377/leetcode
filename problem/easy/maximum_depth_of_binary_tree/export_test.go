package maximum_depth_of_binary_tree

import "testing"

type args struct {
	root *TreeNode
}

type testcase struct {
	name string
	args args
	want int
}

func getTeseCases() []testcase {
	return []testcase{
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
}

func Test_maxDepthRecursively(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthRecursively(tt.args.root); got != tt.want {
				t.Errorf("maxDepthRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthBFS(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthBFS(tt.args.root); got != tt.want {
				t.Errorf("maxDepthRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepthDFS(t *testing.T) {
	for _, tt := range getTeseCases() {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthDFS(tt.args.root); got != tt.want {
				t.Errorf("maxDepthRecursively() = %v, want %v", got, tt.want)
			}
		})
	}
}
