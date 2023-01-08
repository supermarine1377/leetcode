package merge_two_sorted_lists

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type args struct {
	list1 *ListNode
	list2 *ListNode
}

type testcase struct {
	name string
	args args
	want *ListNode
}

func getTestCases() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
				list2: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "2nd",
			args: args{
				list1: &ListNode{
					Val: 1,
				},
				list2: nil,
			},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			name: "3rd",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
			},
		},
		{
			name: "4th",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
	}
}

func Test_mergeTwoLists(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}

func Test_mergeTwoListsRecurcively(t *testing.T) {
	for _, tt := range getTestCases() {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoListsRecursively(tt.args.list1, tt.args.list2)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}
