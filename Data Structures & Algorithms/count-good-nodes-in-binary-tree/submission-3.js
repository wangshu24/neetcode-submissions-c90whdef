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
     * @return {number}
     */
    goodNodes(root) {
        let res = []
        this.dfs(root, Number.MIN_SAFE_INTEGER, res)
        return res.length
    }

    dfs(root, max, res){
        if (!root) {return}

        if (max <= root.val) {res.push(root.val)}
        max = Math.max(max, root.val)

        this.dfs(root.left, max, res)
        this.dfs(root.right, max, res)
    }
}
