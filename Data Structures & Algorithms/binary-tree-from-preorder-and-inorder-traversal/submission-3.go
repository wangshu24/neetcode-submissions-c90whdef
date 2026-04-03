/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	indices := make(map[int]int)

	for index, val := range inorder {
		indices[val] = index
	}

	var res *TreeNode

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(pres, pree, ins, ine int) *TreeNode {
		if pres > pree || ins > ine {return nil}
		root := &TreeNode{
			Val: preorder[pres],
		}

		mid := indices[preorder[pres]]
		lsize := mid - ins
		left := dfs(pres+1, pres+lsize, ins, mid)
		right := dfs(pres+lsize+1, pree, mid+1, ine)
		root.Left = left
		root.Right = right
		return root
	}

	res = dfs(0, len(preorder)-1, 0, len(inorder)-1)

	return res
}
