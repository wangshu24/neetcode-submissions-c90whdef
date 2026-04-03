/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	left := inorderTraversal(root.Left)
	res = append(res, left...)
	res = append(res, root.Val)
	right := inorderTraversal(root.Right)
	res = append(res, right...)
	return res
}
