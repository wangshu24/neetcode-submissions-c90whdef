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
        let res = []
        this.bfs(root, 0, [],res)

        return res
    }

    bfs(node, level, subset, res) {
        if (!node) {return} 

        if (res.length <= level) {res.push([])}
        subset.push(node.val)

        res[level].push(subset)

        this.bfs(node.left,level+1, [], res)
        this.bfs(node.right, level+1, [], res)

    }

}
