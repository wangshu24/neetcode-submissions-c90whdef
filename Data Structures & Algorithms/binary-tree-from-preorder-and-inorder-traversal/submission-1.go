import ("slices")
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 { return nil }

	res := &TreeNode {
		Val: preorder[0],
	}

	idx := slices.Index(inorder, preorder[0])
	if idx == -1 {
		return nil
	}
	res.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	res.Right = buildTree(preorder[1+idx:], inorder[idx+1:])

	return res
}