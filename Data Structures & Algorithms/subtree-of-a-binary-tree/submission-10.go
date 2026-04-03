/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(root, subroot *TreeNode) bool {
	if root == nil && subroot == nil {
		return true
	}
	if root != nil && subroot != nil && root.Val == subroot.Val {
		return isSameTree(root.Left, subroot.Left) && isSameTree(root.Right, subroot.Right)
	} else {
		return false
	}
}
