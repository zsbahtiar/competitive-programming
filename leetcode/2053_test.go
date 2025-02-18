package leetcode

import "testing"

func TestKthDistinctStringInAnArray(t *testing.T) {
	var cases = []struct {
		input  []string
		k      int
		output string
	}{
		{
			input:  []string{"d", "b", "c", "b", "c", "a"},
			k:      2,
			output: "a",
		},
		{
			input:  []string{"aaa", "aa", "a"},
			k:      1,
			output: "aaa",
		},
		{
			input:  []string{"a", "b", "a"},
			k:      3,
			output: "",
		},
	}

	for _, c := range cases {
		got := KthDistinctStringInAnArray(c.input, c.k)
		if got != c.output {
			t.Logf("input: %v, k: %d, got: %s, expected: %s", c.input, c.k, got, c.output)
			t.Fail()
		}
	}
}
