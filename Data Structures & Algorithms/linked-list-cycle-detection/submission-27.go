/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	
    visit := make(map[*ListNode]bool)

	for {
		_, there := visit[head]
		if !there {
			visit[head] = true
			if head.Next == nil {
				return false
			}
			head = head.Next
		} else {
			return true
		}
	}

	return false
}
