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
    dfs(root1, root2){
        if (!root1 && !root2) {return true}
        if (!root1 || !root2) {return false}

        if (root1.val === root2.val) {
            return this.dfs(root1.left, root2.left) && this.dfs(root1.right, root2.right)
        }
    }
    /**
     * @param {TreeNode} root
     * @param {TreeNode} subRoot
     * @return {boolean}
     */
    isSubtree(root, subRoot) {
        if (!root) {return false}
        if (!subRoot) {return true}

        if (this.dfs(root, subRoot)) {return true}

        return this.isSubtree(root.left, subRoot) || this.isSubtree(root.right, subRoot)
    }
}
