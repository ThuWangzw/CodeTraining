class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        ListNode last_p = null;
        ListNode p = head;
        while (p != null) {
            if(p == head){
                p = head.next;
                last_p = head;
                continue;
            }
            if(p.val == last_p.val){
                last_p.next = p.next;
            }
            else {
                last_p = p;
            }
            p = p.next;
        }
        return head;
    }
}