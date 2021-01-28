class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        ListNode cur = head;
        ListNode last = null;
        while ((cur!=null) && (cur.next != null)) {
            boolean eq = false;
            while ((cur.next != null) && (cur.val == cur.next.val)) {
                eq = true;
                cur.next = cur.next.next;
            }
            if(eq) {
                if(last == null) {
                    head = cur.next;
                    cur = cur.next;
                }
                else {
                    last.next = cur.next;
                    cur = cur.next;
                }
            }
            else {
                last = cur;
                cur = cur.next;
            }
        }
        return head;
    }
}