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
     * @param {ListNode} l1
     * @param {ListNode} l2
     * @return {ListNode}
     */
    addTwoNumbers(l1, l2) {
        const res = new ListNode()
        let copy = res
        let carry = 0

        while(l1 || l2 || carry) {
            let val1 = l1 ? l1.val : 0
            let val2 = l2 ? l2.val : 0
            let newNode = new ListNode()
            let newVal = (val1 + val2)

            if (carry > 0) {
                if (newVal > 8) {
                    newNode.val = (newVal + carry) % 10 
                } else {
                    newNode.val = (newVal + carry) % 10
                    --carry
                }
            } else {
                if (newVal > 9) {
                    newNode.val = newVal % 10
                    carry++
                } else {
                    newNode.val = newVal
                }
            }
            console.log(newNode, carry)

            copy.next = newNode
            copy = copy.next
            l1 = l1 ? l1.next : null
            l2 = l2 ? l2.next : null
        }

        return res.next
    }
}
