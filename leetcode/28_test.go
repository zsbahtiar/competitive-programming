package leetcode

import "testing"

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	var cases = []struct {
		name  string
		input struct {
			haystack string
			needle   string
		}
		want int
	}{
		{
			name: "Example 1: First occurrence at beginning",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "Example 2: Needle not found",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "Single character match",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			name: "Single character no match",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "a",
				needle:   "b",
			},
			want: -1,
		},
		{
			name: "Needle longer than haystack",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "abc",
				needle:   "abcd",
			},
			want: -1,
		},
		{
			name: "Multiple occurrences, return first",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "aaaaa",
				needle:   "aa",
			},
			want: 0,
		},
		{
			name: "Occurrence in middle",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			name: "Empty needle",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "hello",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Empty haystack and needle",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
		{
			name: "Empty haystack with non-empty needle",
			input: struct {
				haystack string
				needle   string
			}{
				haystack: "",
				needle:   "a",
			},
			want: -1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FindTheIndexOfTheFirstOccurrenceInAString(tc.input.haystack, tc.input.needle)
			if got != tc.want {
				t.Errorf("FindTheIndexOfTheFirstOccurrenceInAString(%q, %q) = %d; want %d",
					tc.input.haystack, tc.input.needle, got, tc.want)
			}
		})
	}
}
