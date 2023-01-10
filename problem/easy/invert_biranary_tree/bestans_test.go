package invert_biranary_tree

import (
	"reflect"
	"testing"
)

func Test_invertTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "1st",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		{
			name: "2nd",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 8,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 7,
					},
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "3rd",
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := invertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
