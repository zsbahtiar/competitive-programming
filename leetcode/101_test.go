package leetcode

import "testing"

func TestSymmetricTree(t *testing.T) {
	var cases = []struct {
		input  *TreeNode
		output bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			output: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			output: false,
		},
	}

	for _, c := range cases {
		got := SymmetricTree(c.input)
		if got != c.output {
			t.Fail()
		}
	}
}
