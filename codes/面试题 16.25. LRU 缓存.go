type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type LRUCache struct {
	head     *Node
	cacheMap map[int]*Node
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     nil,
		cacheMap: make(map[int]*Node),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		if node == this.head {
			return node.value
		}
		node.left.right = node.right
		node.right.left = node.left
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cacheMap[key]; ok {
		if node == this.head {
			node.value = value
			return
		}
		node.left.right = node.right
		node.right.left = node.left
		this.moveToHead(node)
		node.value = value
	} else {
		node := &Node{
			key:   key,
			value: value,
		}
		if len(this.cacheMap) < this.capacity {
			this.cacheMap[key] = node
			this.moveToHead(node)
		} else {
			// remove tail
			tail := this.head.left
			tail.left.right = this.head
			tail.right.left = tail.left
			delete(this.cacheMap, tail.key)
			this.cacheMap[key] = node
			this.moveToHead(node)
		}
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	if this.head == nil {
		this.head = node
		node.left = node
		node.right = node
	} else {
		node.right = this.head
		node.left = this.head.left
		node.left.right = node
		node.right.left = node
		this.head = node
	}
}