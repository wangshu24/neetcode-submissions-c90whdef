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
        return this.add(l1, l2, 0)
    }

    add(l1, l2, carry) {
        if (!l1 && !l2 && carry === 0 ) {return null}
        
        let v1 = 0
        if (l1){
            v1 = l1.val
        }
        let v2 = 0
        if (l2){
            v2 = l2.val
        } 

        let newVal = v1 + v2
        if (carry > 0) {
            newVal++
            carry--
        }
        if (newVal > 9) {
            carry++
        }
        let newNode = new ListNode()
        newNode.val = newVal%10
        l1 = l1 ? l1.next : null
        l2 = l2 ? l2.next : null
        newNode.next = this.add(l1, l2, carry)
        return newNode
    }
}
