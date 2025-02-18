package leetcode

import "testing"

func TestCheckIfWordEqualsSummationOfTwoWords(t *testing.T) {
	var cases = []struct {
		input struct {
			firstWord  string
			secondWord string
			targetWord string
		}
		output bool
	}{
		{
			input: struct {
				firstWord  string
				secondWord string
				targetWord string
			}{
				firstWord:  "acb",
				secondWord: "cba",
				targetWord: "cdb",
			},
			output: true,
		},
		{
			input: struct {
				firstWord  string
				secondWord string
				targetWord string
			}{
				firstWord:  "aaa",
				secondWord: "a",
				targetWord: "aab",
			},
			output: false,
		},
		{
			input: struct {
				firstWord  string
				secondWord string
				targetWord string
			}{
				firstWord:  "aaa",
				secondWord: "a",
				targetWord: "aaaa",
			},
			output: true,
		},
	}

	for _, test := range cases {
		got := CheckIfWordEqualsSummationOfTwoWords(test.input.firstWord, test.input.secondWord, test.input.targetWord)
		if got != test.output {
			t.Fail()
		}
	}
}
