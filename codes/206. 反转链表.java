class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode p = head;
        ListNode last = null;
        ListNode newp = null;
        while (p!=null) {
            newp = p.next;
            p.next = last;
            last = p;
            p = newp;
        }
        return last;
    }
}