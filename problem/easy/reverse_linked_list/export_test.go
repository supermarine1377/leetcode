package reverse_linked_list

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type args struct {
	head *listNode
}

type testcase struct {
	name string
	args args
	want *listNode
}

func getTestCase() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				head: &listNode{
					Val: 1,
					Next: &listNode{
						Val: 2,
						Next: &listNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
			want: &listNode{
				Val: 3,
				Next: &listNode{
					Val: 2,
					Next: &listNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
		{
			name: "2nd",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
}

func Test_reverseList(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			opt := cmp.AllowUnexported(listNode{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}

func Test_reverseListRecurcively(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			got := reveseListRecursively(tt.args.head)
			opt := cmp.AllowUnexported(listNode{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}

func Test_reverseLink_stack(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseLinkList_stack(tt.args.head)
			opt := cmp.AllowUnexported(listNode{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}
