package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMode(t *testing.T) {
	var root *TreeNode
	var got, expected []int

	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	expected = []int{2}
	got = findMode(root)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("findMode() = %v; expected %v", got, expected)
	}

	root = &TreeNode{
		Val: 0,
	}
	expected = []int{0}
	got = findMode(root)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("findMode() = %v; expected %v", got, expected)
	}
}
