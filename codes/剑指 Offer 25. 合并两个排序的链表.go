/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head *ListNode
    var tail *ListNode
    p := l1
    q := l2
    for p!=nil && q!=nil {
        var next *ListNode
        if p.Val>q.Val {
            next = q
        } else {
            next = p
        }
        if head==nil {
            head = next
            tail = next
        } else {
            tail.Next = next
            tail = next
        }
        if next==p {
            p=p.Next
        } else {
            q=q.Next
        }
    }
    for p!=nil {
        if head==nil {
            head = p
            tail = p
        } else {
            tail.Next = p
            tail = p
        }
        p = p.Next
    }
    for q!=nil {
        if head==nil {
            head = q
            tail = q
        } else {
            tail.Next = q
            tail = q
        }
        q = q.Next
    }
    return head
}