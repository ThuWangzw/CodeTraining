/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val:  node.Val,
			Next: node.Next,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	newhead := head.Next
	for node := head; node != nil; node = node.Next {
		if node.Next.Next != nil {
			node.Next, node.Next.Next = node.Next.Next, node.Next.Next.Next
		} else {
			node.Next = nil
		}
	}
	return newhead
}