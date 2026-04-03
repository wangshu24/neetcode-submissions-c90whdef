/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    equal := true 

	var dfsc func(*TreeNode, *TreeNode)
	dfsc = func(p, q *TreeNode) {
		if q == nil && p == nil {
			return
		}
		if q != nil && p != nil && q.Val == p.Val  {
			dfsc(q.Left, p.Left)
			dfsc(q.Right, p.Right)
		} else {
			equal = false
			return
		}
	}
	dfsc(p,q)

	return equal
}
