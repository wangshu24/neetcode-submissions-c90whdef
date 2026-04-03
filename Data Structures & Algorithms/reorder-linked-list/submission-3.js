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
     * @return {void}
     */
    reorderList(head) {
        let slow = head
        let fast = head.next

        while (fast && fast.next) {
            fast = fast.next.next
            slow = slow.next
        }
        let cur = slow.next
        let tail = slow.next = null

        while (cur) {
            const tmp = cur.next
            cur.next = tail
            tail = cur
            cur = tmp
        }

        console.log(tail)
        cur = head
        while(tail && cur) {
            const nextRoot = cur.next
            const nextTail = tail.next
            cur.next = tail
            tail.next = nextRoot
            cur = nextRoot
            tail = nextTail
        }
    }
}
