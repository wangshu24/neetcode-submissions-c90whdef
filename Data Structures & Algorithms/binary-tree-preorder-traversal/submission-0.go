/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{root.Val}
	left := preorderTraversal(root.Left)
	res = append(res, left...)
	right := preorderTraversal(root.Right)
	res = append(res, right...)
	return res
   
}
