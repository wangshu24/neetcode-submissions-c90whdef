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
     * @param {number} k
     * @return {number}
     */
    kthSmallest(root, k) {
        let res = []
        this.dfs(root, res)
       
        return res.sort()[k-1]
    }

    dfs(root,res){
        if(!root) {return}
        res.push(root.val)

        this.dfs(root.left,res)
        this.dfs(root.right, res)
    }
}
