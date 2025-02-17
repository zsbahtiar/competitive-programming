package leetcode

import "testing"

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}},
			},
			output: 3,
		},
		{
			input:  &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			output: 2,
		},
	}

	for _, c := range cases {
		got := MaximumDepthOfBinaryTree(c.input)
		if got != c.output {
			t.Fail()
		}
	}
}
