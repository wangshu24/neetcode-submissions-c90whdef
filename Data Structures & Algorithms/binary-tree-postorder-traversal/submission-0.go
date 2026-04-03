/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	left := postorderTraversal(root.Left)
	res = append(res, left...)
	right := postorderTraversal(root.Right)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}
