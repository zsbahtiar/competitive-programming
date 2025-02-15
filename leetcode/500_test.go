package leetcode

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	var words []string
	var got, expected []string

	words = []string{"Hello", "Alaska", "Dad", "Peace"}
	expected = []string{"Alaska", "Dad"}
	got = findWords(words)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("findWords() = %v; expected %v", got, expected)
	}
}
