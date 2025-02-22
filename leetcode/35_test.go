package leetcode

import "testing"

func TestSearchInsertPositiont(t *testing.T) {
	var cases = []struct {
		input struct {
			nums   []int
			target int
		}
		want int
	}{
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			want: 4,
		},
		{
			input: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 3, 5},
				target: 3,
			},
			want: 1,
		},
	}

	for _, test := range cases {
		got := SearchInsertPositiont(test.input.nums, test.input.target)
		if got != test.want {
			t.Fail()
		}
	}
}
