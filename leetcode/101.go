package leetcode

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

// Example 1
// Input: root = [1,2,2,3,4,4,3]
// Output: true

// Example 2
// Input: root = [1,2,2,null,3,null,3]
// Output: false

// Constraints:

// The number of nodes in the tree is in the range [1, 1000].
// -100 <= Node.val <= 100
type TreeNodeSymetric struct {
	Val   int
	Left  *TreeNodeSymetric
	Right *TreeNodeSymetric
}

func SymmetricTree(root *TreeNodeSymetric) bool {
	symmetric := true
	if root == nil {
		return symmetric
	}

	// process each node by queue, after check remove (FIFO)
	var queue []*TreeNodeSymetric
	queue = append(queue, root.Left, root.Right)

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			symmetric = false
			break
		} else if left.Val != right.Val {
			symmetric = false
			break
		}

		queue = append(queue, left.Left, right.Right)
		queue = append(queue, left.Right, right.Left)
	}

	return symmetric

}
