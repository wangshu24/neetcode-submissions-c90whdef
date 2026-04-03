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
	newCopy := &Node{
		Val: 0,
	}

	// newcopy2 := newCopy
	res := newCopy
	copyHead := head
	for copyHead != nil {
		newNode := &Node{
			Val: copyHead.Val,
		}
		newCopy.Next = newNode
		newCopy = newCopy.Next
		visit[copyHead] = newNode
		copyHead = copyHead.Next
	}

	copyHead2 := head
	for copyHead2 != nil {
		if copyHead2.Random != nil {
			visit[copyHead2].Random = visit[copyHead2.Random]
		} else {
			visit[copyHead2].Random = nil
		}
		copyHead2 = copyHead2.Next
	}

	return res.Next
}
