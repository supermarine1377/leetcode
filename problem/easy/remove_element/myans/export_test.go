package myans

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		Val  int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		{
			name: "1st",
			args: args{
				nums: []int{3, 2, 2, 3},
				Val:  3,
			},
			want:    2,
			wantArr: []int{2, 2, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.Val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			if diff := cmp.Diff(tt.args.nums, tt.wantArr); diff != "" {
				t.Errorf("error: unexpected diff; %s", diff)
			}
		})
	}
}
