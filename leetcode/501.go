package leetcode

// Given the root of a binary search tree (BST) with duplicates, return all the mode(s) (i.e., the most frequently occurred element) in it.
// If the tree has more than one mode, return them in any order.
// Assume a BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than or equal to the node's key.
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
// Both the left and right subtrees must also be binary search trees.
// Input: root = [1,null,2,2]
// Output: [2]
// Example 2:
// Input: root = [0]
// Output: [0]

// Constraints:
// The number of nodes in the tree is in the range [1, 104].
// -105 <= Node.val <= 105
// Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var (
		result []int
		max    int
		count  int
		prev   int
	)
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev == node.Val {
			count++
		} else {
			count = 1
			prev = node.Val
		}
		if count == max {
			result = append(result, node.Val)
		} else if count > max {
			max = count
			result = []int{node.Val}
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}
