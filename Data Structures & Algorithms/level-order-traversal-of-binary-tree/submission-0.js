/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     constructor(val = 0, left = null, right = null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    /**
     * @param {TreeNode} root
     * @return {number[][]}
     */
    levelOrder(root) {
        let res = this.bfs(root, 0, [])
        return res
    }

    bfs(root, level, res) {
        if (!root) {return []}

       if (res.length <= level) {res.push([])}

       res[level].push(root.val)

       this.bfs(root.left, level+1, res)
       this.bfs(root.right, level+1, res)

       return res
    }
}
