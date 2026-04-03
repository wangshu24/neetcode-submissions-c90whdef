/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

class Solution {
    /**
     * @param {ListNode} head
     * @return {boolean}
     */
    hasCycle(head) {
        let visited = []
        let index = -1
        let temp = -1

        while (head && index === -1) {
            temp++
            if (!visited.includes(head.val)) {
                visited.push(head.val)
            } else {
                index = temp
            }
            head = head.next
        }

        if (index !== -1) {return true}
        return false

    }
}
