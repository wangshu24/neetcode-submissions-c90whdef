/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := math.MinInt32

	var dfs func(*TreeNode) int
	dfs = func(r *TreeNode) int {
		if r == nil {return 0}
	
		leftMax := dfs(r.Left)

		rightMax := dfs(r.Right)

		rMax := max(r.Val, r.Val + leftMax, r.Val + rightMax)
		
		res = max(res,r.Val, r.Val + leftMax + rightMax, r.Val + leftMax, r.Val + rightMax)
		return rMax
	}

	_ = dfs(root)
	return res
}