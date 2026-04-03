/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nNode := head
	for i:=0; i < n; i++ {
		nNode = nNode.Next
	} 
	if nNode == nil {
		head = head.Next
		return head
	}

	fmt.Println(nNode)
	temp := head
	for nNode.Next != nil {
		temp = temp.Next
		nNode = nNode.Next
	}
	fmt.Println(nNode, temp)
	after := temp.Next.Next
	remove := temp.Next
	remove.Next = nil
	temp.Next = after
	
	return head
}
