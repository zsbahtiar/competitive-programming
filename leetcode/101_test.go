package leetcode

import "testing"

func TestSymmetricTree(t *testing.T) {
	var cases = []struct {
		input  *TreeNodeSymetric
		output bool
	}{
		{
			input: &TreeNodeSymetric{
				Val: 1,
				Left: &TreeNodeSymetric{
					Val: 2,
					Left: &TreeNodeSymetric{
						Val: 3,
					},
					Right: &TreeNodeSymetric{
						Val: 4,
					},
				},
				Right: &TreeNodeSymetric{
					Val: 2,
					Left: &TreeNodeSymetric{
						Val: 4,
					},
					Right: &TreeNodeSymetric{
						Val: 3,
					},
				},
			},
			output: true,
		},
		{
			input: &TreeNodeSymetric{
				Val: 1,
				Left: &TreeNodeSymetric{
					Val: 2,
					Right: &TreeNodeSymetric{
						Val: 3,
					},
				},
				Right: &TreeNodeSymetric{
					Val:   2,
					Right: &TreeNodeSymetric{Val: 3},
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
