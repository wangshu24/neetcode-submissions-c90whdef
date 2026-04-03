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

        let res =  [] 
        let lvl = this.bfs(root,0,[])
        lvl.forEach(e => res.push(e.pop()))
        return res
    }

    bfs(root, level, lvlorder) {
        if (!root) {return []}

        if (lvlorder.length <= level) {lvlorder.push([])}
        
        lvlorder[level].push(root.val)
        this.bfs(root.left, level+1, lvlorder)
        this.bfs(root.right, level+1, lvlorder)
        return lvlorder
    }
}
