package leetcode

import "testing"

func TestValidParentheses(t *testing.T) {
	var cases = []struct {
		input  string
		output bool
	}{
		{
			input:  "",
			output: true,
		},
		{
			input:  "()",
			output: true,
		},
		{
			input:  "()[]{}",
			output: true,
		},
		{
			input:  "(]",
			output: false,
		},
		{
			input:  "([])",
			output: true,
		},
	}

	for _, c := range cases {
		got := ValidParentheses(c.input)
		if got != c.output {
			t.Logf("failed: got %t want %t in: %s\n", got, c.output, c.input)
			t.Fail()
		}
	}

}
