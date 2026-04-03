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
     * @param {TreeNode} root1
     * @param {TreeNode} root2
     * @return {boolean}
     */
    dfs(root1, root2) {
        if (!root1 && !root2) {
            return true
        }

        if (!root1) {return false}
        if (!root2) {return false}

        if (root1.val !== root2.val) {return false}

        const left = this.dfs(root1.left, root2.left)
        const right = this.dfs(root1.right, root2.right)
        return left && right
    }

    /**
     * @param {TreeNode} p
     * @param {TreeNode} q
     * @return {boolean}
     */
    isSameTree(p, q) {
        return this.dfs(p, q)
    }
}
