/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var res int
	count := 0
	var dfs func(*TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		} else if count == k {
			return false
		}
		
		if found := dfs(root.Left); found {return true}
		count++
		if count == k  {
			res = root.Val
			return false
		}
		if found := dfs(root.Right); found {return true}
		return false
	}

	_ = dfs(root)

	return res
}