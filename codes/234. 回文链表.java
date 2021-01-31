class Solution {
    ListNode reverseList(ListNode head) {
        if(head==null || head.next ==null) return head;
        ListNode p = head.next;
        ListNode q = head;
        head.next = null;
        while (p != null) {
            ListNode tmp = p.next;
            p.next = q;
            q = p;
            p = tmp;
        }
        return q;
    }

    public boolean isPalindrome(ListNode head) {
        if(head==null || head.next==null) return true;
        // 1. find mid point
        ListNode fast = head;
        ListNode slow = head;
        while (fast.next != null && fast.next.next != null) {
            fast = fast.next.next;
            slow = slow.next;
        }
        if(fast.next!=null) fast = fast.next;

        // 2. reverse right
        fast = reverseList(slow.next);

        // 3. compare list
        slow = head;
        while (fast!=null) {
            if(slow.val != fast.val) return false;
            slow = slow.next;
            fast = fast.next;
        }
        return true;
    }
}