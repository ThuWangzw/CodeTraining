// 哈希表+双向链表
type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}

type LRUCache struct {
	head     *Node
	length   int
	capacity int
	keyMap   map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyMap:   make(map[int]*Node),
		capacity: capacity,
	}
}

func (this *LRUCache) putHead(val *Node) {
	if val != this.head && val != this.head.left {
		val.left.right = val.right
		val.right.left = val.left
		tail := this.head.left
		val.left = tail
		val.right = this.head
		tail.right = val
		this.head.left = val
	}
	this.head = val
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.keyMap[key]
	if !ok {
		return -1
	} else {
		this.putHead(val)
		return val.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.keyMap[key]; ok {
		val.val = value
		this.putHead(val)
	} else if this.length < this.capacity {
		node := &Node{
			key: key,
			val: value,
		}
		if this.length > 0 {
			node.left = this.head.left
			node.right = this.head
		} else {
			node.left = node
			node.right = node
			this.head = node
		}
		this.head.left.right = node
		this.head.left = node
		this.head = node
		this.keyMap[key] = node
		this.length++
	} else {
		delete(this.keyMap, this.head.left.key)
		this.head.left.key = key
		this.head.left.val = value
		this.head = this.head.left
		this.keyMap[key] = this.head
	}
}