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
     * @param {number} p
     * 
     */
    // dfs(root, p, q) {
    //     if (!root) {return false}
    //     if (root.val)

    // }

    /**
     * @param {TreeNode} root
     * @param {TreeNode} p
     * @param {TreeNode} q
     * @return {TreeNode}
     */
    lowestCommonAncestor(root, p, q) {
       if (!root || !q || !p) {return null}

       if (Math.max(q.val, p.val) < root.val) {return this.lowestCommonAncestor(root.left, p, q)} 
       else if (Math.min(q.val, p.val) > root.val) {
        return this.lowestCommonAncestor(root.right, p, q)
       } else {return root}
    } 
}
