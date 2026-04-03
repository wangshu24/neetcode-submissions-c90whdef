/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	
	var prev *ListNode
	temp2 := slow.Next
	slow.Next = nil
	for temp2 != nil {
		next := temp2.Next
		temp2.Next = prev
		prev = temp2
		temp2 = next
	}

	// fmt.Println(prev, slow)
	
	copyhead := head
	for prev != nil && copyhead != nil {
		nextL := copyhead.Next
		nextR := prev.Next
		copyhead.Next = prev
		prev.Next = nextL
		prev = nextR
		copyhead = nextL  
	}
}
