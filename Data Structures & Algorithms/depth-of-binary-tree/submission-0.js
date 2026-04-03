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
    maxDepth(root) {
        if (!root) {
            return 0
        }

        // if (!root.left) {
        //     return this.maxDepth(root.right)
        // }
        // if (!root.right) {
        //     return this.maxDepth(root.left)
        // }

        let res = 1
        let depL = this.maxDepth(root.left)
        let depR = this.maxDepth(root.right)
        if (depL > depR) {
            res = res+ depL
        } else {res = res + depR}
        return res
    }
}
