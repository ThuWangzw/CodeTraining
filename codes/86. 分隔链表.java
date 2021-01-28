class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode left_head = new ListNode(-1, null);
        ListNode left_tail = left_head;
        ListNode right_head = new ListNode(-1, null);
        ListNode right_tail = right_head;
        ListNode p = head;
        while (p!=null) {
            if(p.val < x) {
                left_tail.next = p;
                left_tail = left_tail.next;
            }
            else {
                right_tail.next = p;
                right_tail = right_tail.next;
            }
            p = p.next;
        }
        left_tail.next = right_head.next;
        right_tail.next = null;
        return left_head.next;
    }
}