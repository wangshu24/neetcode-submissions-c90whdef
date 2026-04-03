/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    visit := make(map[*Node]*Node)
	newHead := &Node {
		Val: 0,
		Next: nil,
		Random: nil,
	}
	copyNewHead := newHead
	oldHead := head
	for head != nil {
		newNode := Node{
			Val: head.Val,
			Next: nil,
		}
		copyNewHead.Next = &newNode
		visit[head] = &newNode
		copyNewHead = copyNewHead.Next
		head = head.Next
	}
	copyNewHead2 := newHead.Next
	fmt.Println(copyNewHead2)
	for oldHead != nil {
		if oldHead.Random != nil {
			visit[oldHead].Random = visit[oldHead.Random] 
		} else {
			copyNewHead2.Random = nil
		}
		copyNewHead2 = copyNewHead2.Next
		oldHead = oldHead.Next
	} 

	return newHead.Next
}
