package reverse_string

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1st",
			args: args{
				s: []byte{10, 20, 30, 40},
			},
			want: []byte{40, 30, 20, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.s
			reverseString(got)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("error :unexpeced diff %s", diff)
			}
		})
	}
}
