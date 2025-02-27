package leetcode

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	var tests = []struct {
		a, b, want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, test := range tests {
		got := AddBinary(test.a, test.b)
		if got != test.want {
			t.Errorf("AddBinary(%q, %q) = %q; want %q", test.a, test.b, got, test.want)
			t.Fail()
		}
	}
}
