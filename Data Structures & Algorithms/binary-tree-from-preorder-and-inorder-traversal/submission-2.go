/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    index := make(map[int]int)
	for ind,val := range inorder {
		index[val] = ind
	}

	var res *TreeNode

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(prestart, preend, instart, inend int) *TreeNode {
		if preend < prestart || inend < instart {
			return nil
		}
		root := &TreeNode{
			Val: preorder[prestart],
		}

		mid := index[preorder[prestart]]
		leftsize := mid - instart
		left := dfs(prestart+1, prestart + leftsize, instart,mid-1)
		right := dfs(prestart+leftsize+1, preend,mid+1, inend)
		root.Left = left
		root.Right = right
		return root
	}
	
	res = dfs(0, len(preorder)-1,0, len(inorder)-1)
	return res
}