package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var cases = []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"flower", "flow", "flight"},
			output: "fl",
		},
		{
			input:  []string{"dog", "racecar", "car"},
			output: "",
		},
	}
	for _, c := range cases {
		x := longestCommonPrefix(c.input)
		if x != c.output {
			t.Fail()
		}
	}
}
