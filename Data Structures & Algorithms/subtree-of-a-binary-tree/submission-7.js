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
    same(root1, root2) {
        if (!root2 && !root1) {return true}
        
        if (root1 && root2 && root1.val === root2.val) {
            let left = this.same(root1.left, root2.left)
            let right = this.same(root1.right, root2.right)

            return left && right
        } 
        return false
    }

    // dfs(root1, root2) {
    //     if (!root1 && !root2) {return true}
    //     if (!root1 || !root2)  {return false}

    //     if (root1.val === root2.val) {
    //         //dfs same
    //         let sameLeft = this.same(root1.left, root2.left)
    //         let sameRight = this.same(root1.right, root2.right)
    //         return sameLeft && sameRight
    //     } else {
    //         let left = this.dfs(root1.left, root2)
    //         let right = this.dfs(root1.right, root2)
    //         return left || right
    //     }
    // }
    /**
     * @param {TreeNode} root
     * @param {TreeNode} subRoot
     * @return {boolean}
     */
    isSubtree(root, subRoot) {
        if (!subRoot) {return true}
        if (!root) {return false}

        if (this.same(root, subRoot)) {return true}
        return this.isSubtree(root.left, subRoot) || this.isSubtree(root.right, subRoot) 
    }
}
