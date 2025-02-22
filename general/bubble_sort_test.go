package general

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		arg  []int
		want []int
	}{
		{
			arg:  []int{32, 27, 43, 3, 9, 82, 10},
			want: []int{3, 9, 10, 27, 32, 43, 82},
		},
		{
			arg:  []int{11, 6, 3, 24, 46, 22, 7},
			want: []int{3, 6, 7, 11, 22, 24, 46},
		},
	}

	for _, test := range tests {
		got := BubbleSort(test.arg)
		if !reflect.DeepEqual(&got, &test.want) {
			t.Logf("arg: %v, want: %v, got: %v", test.arg, test.want, got)
			t.Fail()
		}
	}
}
