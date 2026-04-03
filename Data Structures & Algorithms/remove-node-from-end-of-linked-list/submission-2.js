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
     * @param {number} n
     * @return {ListNode}
     */
    removeNthFromEnd(head, n) {
        let tmp = head
        let count = 0
        while (tmp!== null) {
            tmp = tmp.next
            ++count
        }

        if (n > count) {return head}
        let tar = count - n
        console.log(count, tar)
        let cur = head
        for (let i=0; i < count; i++) {
            if (i === tar && i === 0 ) {
                head = cur.next
                cur = cur.next
            } else if (i+1 !== tar) {
                cur = cur.next
            } else {
                const after = cur.next.next
                cur.next.next = null
                cur.next = after
                break
            }
        }

        return head
    }
}
