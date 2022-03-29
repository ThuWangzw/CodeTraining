type Node struct {
	key   int
	value int
	left  *Node
	right *Node
}

type LRUCache struct {
	cacheMap  map[int]*Node
	cacheHead *Node
	capacity  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheMap:  make(map[int]*Node),
		cacheHead: nil,
		capacity:  capacity,
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	if this.cacheHead == node {
		return
	}
	// first, remove node from linkedlist
	node.left.right = node.right
	node.right.left = node.left
	// second, add to head
	node.left = this.cacheHead.left
	node.right = this.cacheHead
	node.left.right = node
	node.right.left = node
	this.cacheHead = node
}

func (this *LRUCache) newNode(key int, value int) {
	if this.cacheHead == nil {
		this.cacheHead = &Node{
			key:   key,
			value: value,
		}
		this.cacheHead.left = this.cacheHead
		this.cacheHead.right = this.cacheHead
		this.cacheMap[this.cacheHead.key] = this.cacheHead
		return
	}
	node := &Node{
		key:   key,
		value: value,
		left:  this.cacheHead.left,
		right: this.cacheHead,
	}
	node.left.right = node
	node.right.left = node
	this.cacheHead = node
	this.cacheMap[node.key] = node
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cacheMap[key]; ok {
		this.moveToHead(node)
		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cacheMap[key]; ok {
		// in cache
		node.value = value
		this.moveToHead(node)
	} else if len(this.cacheMap) < this.capacity {
		// cache not full, add to head
		this.newNode(key, value)
	} else {
		// cache full, delete tail
		tail := this.cacheHead.left
		delete(this.cacheMap, tail.key)
		tail.left.right = tail.right
		tail.right.left = tail.left
		this.newNode(key, value)
	}
}