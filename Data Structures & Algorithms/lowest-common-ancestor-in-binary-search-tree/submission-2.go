/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    res := root

	var recur func(*TreeNode)
	recur = func(root *TreeNode) {
		if root == nil {
			return 
		}
		if root.Val < q.Val && root.Val < p.Val {
			res = root.Right
			recur(root.Right)
		} else if root.Val > q.Val && root.Val > p.Val {
			res = root.Left
			recur(root.Left)
		} else {
			res = root
			return
		}
	}
	recur(root)

	return res
}
