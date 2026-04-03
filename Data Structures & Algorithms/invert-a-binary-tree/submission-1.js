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
     * @return {TreeNode}
     */
    invertTree(root) {
        if (!root || (!root.left && !root.right)) {
            return root
        }

        // if (!root.left) {
        //     root.left = root.right
        //     root.right = null
        //     return root
        // }

        // if (!root.right) {
        //     root.right = root.left
        //     root.left = null
        //     return root
        // }
        let temp = root.left
        root.left = root.right
        root.right = temp
        this.invertTree(root.right)
        this.invertTree(root.left)
        return root
    }
}
