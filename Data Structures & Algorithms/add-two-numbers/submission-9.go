/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode {
		Val :0,
	}
	remain, copyres  := 0, res
	for l1 != nil || l2 != nil || remain != 0 {
		var total int
		if l1 != nil {
			total += l1.Val
		}
		if l2 != nil {
			total += l2.Val
		}
		total += remain
		newVal := total % 10
		remain = total / 10
		newNode := &ListNode {
			Val : newVal,
		}
		copyres.Next = newNode
		copyres  = copyres.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return res.Next
}
