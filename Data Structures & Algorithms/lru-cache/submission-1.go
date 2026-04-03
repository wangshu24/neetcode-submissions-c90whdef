type LRUCache struct {
    cap int
	cache map[int]*Node
	head,tail *Node
}

type Node struct {
	key, val int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache {
		cap: capacity,
		cache: make(map[int]*Node),
		head: &Node{},
		tail: &Node{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *Node) {
	prev, next := this.tail.prev, this.tail
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev 	
}

func (this *LRUCache) Get(key int) int {
	if val, there := this.cache[key]; there {
		this.remove(val)
		this.insert(val)
		return val.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, there := this.cache[key]; there {
		this.remove(val)
		this.insert(val)
		this.cache[key].val = value
		return
	}
	newNode := &Node{
		key: key,
		val: value,
		next : &Node{},
		prev: &Node{},
	}
	if len(this.cache) == this.cap {
		least := this.head.next
		this.remove(least)
		delete(this.cache, least.key)
	}
	this.insert(newNode)
	this.cache[key] = newNode
	return
}
