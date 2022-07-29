/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		node := &Node{
			Val: x,
		}
		node.Next = node
		return node
	}
	originHead := aNode
	for true {
		if aNode.Val != aNode.Next.Val {
			break
		}
		aNode = aNode.Next
		if originHead == aNode {
			break
		}
	}
	var minNode, minNodeLast *Node
	var maxNode, maxNodeLast *Node
	cur, last := aNode.Next, aNode
	for true {
		if cur.Val > x && (minNode == nil || minNode.Val > cur.Val) {
			minNode, minNodeLast = cur, last
		}
		if maxNode == nil || maxNode.Val <= cur.Val {
			maxNode, maxNodeLast = cur, last
		}
		cur, last = cur.Next, cur
		if last == aNode {
			break
		}
	}
	newnode := &Node{Val: x}
	if minNode != nil {
		minNodeLast.Next = newnode
		newnode.Next = minNode
	} else {
		maxNode, maxNodeLast = maxNode.Next, maxNodeLast.Next
		maxNodeLast.Next = newnode
		newnode.Next = maxNode
	}

	return originHead
}