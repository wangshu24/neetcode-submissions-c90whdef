/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
	
	var bfs func(*TreeNode, int) 
	bfs = func(root *TreeNode, level int) {
		if root == nil {
			return 
		}
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)
		bfs(root.Left, level+1)
		bfs(root.Right, level+1)
		return
	}
	
	bfs(root, 0)

	return res 
}
