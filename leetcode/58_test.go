package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	var cases = []struct {
		arg  string
		want int
	}{
		{
			arg:  "Hello World",
			want: 5,
		},
		{
			arg:  "   fly me   to   the moon  ",
			want: 4,
		},
		{
			arg:  "luffy is still joyboy",
			want: 6,
		},
	}

	for _, c := range cases {
		got := LengthOfLastWord(c.arg)
		if got != c.want {
			t.Logf("given: %s, got: %d, want %d", c.arg, got, c.want)
			t.Fail()
		}
	}
}
