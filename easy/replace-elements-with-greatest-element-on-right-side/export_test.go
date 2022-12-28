package replace_elements_with_greatest_element_on_right_side

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "1st",
		args: args{
			arr: []int{17, 18, 5, 4, 6, 1},
		},
		want: []int{18, 6, 6, 6, 1, -1},
	},
	{
		name: "2nd",
		args: args{
			arr: []int{400},
		},
		want: []int{-1},
	},
}

func Test_replaceElements02(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElementsO2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElementsO2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceElements01(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElementsO1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElementsO1() = %v, want %v", got, tt.want)
			}
		})
	}
}
