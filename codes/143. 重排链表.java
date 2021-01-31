/*
    该题目是三个技巧的合并：快慢指针，反转链表和链表合并
 */
class Solution {
    ListNode reverseList(ListNode head) {
        if(head==null || head.next==null) return head;
        ListNode p = head.next;
        ListNode q = head;
        head.next = null;
        while (p!=null) {
            ListNode tmp = p.next;
            p.next = q;
            q = p;
            p = tmp;
        }
        return q;
    }

    void mergeList(ListNode left, ListNode right) {
        ListNode p1 = left;
        ListNode p2 = right;
        while (p2 != null) {
            ListNode tmp = p2.next;
            p2.next = p1.next;
            p1.next = p2;
            p1 = p2.next;
            p2 = tmp;
        }
    }

    public void reorderList(ListNode head) {
        if(head==null || head.next == null) return;
        // 1. find mid point
        ListNode slow = head;
        ListNode fast = head;
        while (fast.next!=null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        if(fast.next != null) fast = fast.next;
        // 2. reverse list
        fast = reverseList(slow.next);
        slow.next = null;
        // 3. merge list
        mergeList(head, fast);
    }
}