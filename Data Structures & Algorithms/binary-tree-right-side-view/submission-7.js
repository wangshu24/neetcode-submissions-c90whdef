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
     * @return {number[]}
     */
    rightSideView(root) {
        let res = []
        this.bfs(root, 0, res)
        let final = res.map(e => e.pop())
        return final
    }

    bfs(root, level, lvlOrder) {
        if (!root) {return}
        
        if (lvlOrder.length <= level) {lvlOrder.push([])}

        lvlOrder[level].push(root.val)

        this.bfs(root.left, level+1, lvlOrder)
        this.bfs(root.right, level+1, lvlOrder) 
    }
}
