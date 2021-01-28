/*
dummy node可以让代码简洁很多
 */
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode p = l1;
        ListNode q = l2;
        ListNode dummy_head = new ListNode(-1, null);
        ListNode tail = dummy_head;
        while ((p!=null) && (q!=null)) {
            if(p.val > q.val) {
                tail.next = q;
                q = q.next;
            }
            else {
                tail.next = p;
                p = p.next;
            }
            tail = tail.next;
        }
        if(p!=null) tail.next = p;
        if(q!=null) tail.next = q;
        return dummy_head.next;
    }
}