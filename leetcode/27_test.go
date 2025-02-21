package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	var cases = []struct {
		input struct {
			nums []int
			val  int
		}
		want int
	}{
		{
			input: struct {
				nums []int
				val  int
			}{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			input: struct {
				nums []int
				val  int
			}{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
	}

	for _, test := range cases {
		got := RemoveElement(test.input.nums, test.input.val)
		if got != test.want {
			x := []int{4, 2, 3}
			t.Log(x[:1])
			t.Fail()
		}
	}
}
