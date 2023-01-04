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
					val: 1,
					next: &listNode{
						val: 2,
						next: &listNode{
							val:  3,
							next: nil,
						},
					},
				},
			},
			want: &listNode{
				val: 3,
				next: &listNode{
					val: 2,
					next: &listNode{
						val:  1,
						next: nil,
					},
				},
			},
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
