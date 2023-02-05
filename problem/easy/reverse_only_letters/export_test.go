package reverse_only_letters

import "testing"

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1st",
			args: args{
				s: "abcd",
			},
			want: "dcba",
		},
		{
			name: "2nd",
			args: args{
				s: "ab-cd",
			},
			want: "dc-ba",
		},
		{
			name: "3rd",
			args: args{
				s: "12-34",
			},
			want: "12-34",
		},
		{
			name: "4th",
			args: args{
				s: "Test1ng-Leet=code-Q!",
			},
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseOnlyLetters(tt.args.s); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
