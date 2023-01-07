package remove_all_adjacent_duplicates_in_string

import "testing"

type args struct {
	s string
}

type testcase struct {
	name string
	args args
	want string
}

func getTestCase() []testcase {
	return []testcase{
		{
			name: "1st",
			args: args{
				s: "abbaca",
			},
			want: "ca",
		},
		{
			name: "2nd",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}
}

func Test_removeDuplicates_ON3(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates_ON3(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates_ON1(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates_ON1(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates_ON1_fastest(t *testing.T) {
	for _, tt := range getTestCase() {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates_ON1_fastest(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
