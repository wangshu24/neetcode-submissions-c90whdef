type Node struct {
	key, val int
	prev, next *Node
}

type LRUCache struct {
    nodes map[int]*Node
	head, tail *Node
	capt int
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{0,0, nil, nil}, &Node{0,0,nil,nil}
	head.prev = tail
	tail.next = head
	return LRUCache{
		nodes: make(map[int]*Node),
		head: head,
		tail: tail,
		capt: capacity,
	}
}

func remove(n *Node) {
	n.next.prev = n.prev
	n.prev.next = n.next
}

func (this *LRUCache) moveToHead(n *Node) {
	n.next = this.head
	n.prev = this.head.prev
	this.head.prev.next = n
	this.head.prev = n
}

func (this *LRUCache) Get(key int) int {
    if node, there := this.nodes[key]; there {
		remove(node)
		this.moveToHead(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, there := this.nodes[key]; there {
		remove(node)
		node.val = value
		this.moveToHead(node)
	} else {
		if len(this.nodes) == this.capt {
			lru := this.tail.next
			remove(lru)
			delete(this.nodes, lru.key) 
		}
		newNode := &Node{key,value, nil,nil}
		this.moveToHead(newNode)
		this.nodes[key] = newNode
	}
}
