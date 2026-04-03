/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode {
		Val: 0,
	}
	
	copyres, remain := res, 0

	for l1 != nil || l2 != nil || remain != 0 {
		var newVal int
		if l1 != nil {
			newVal += l1.Val
		} 
		if l2 != nil {
			newVal += l2.Val
		}
		total := newVal + remain
		newVal = total % 10
		remain = total / 10
		newNode := &ListNode {
			Val: newVal,
		}
		copyres.Next = newNode
		copyres = copyres.Next
		if l1 != nil { l1 = l1.Next}
		if l2 != nil {l2 = l2.Next}
	}

	return res.Next
}
