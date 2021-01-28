class Solution {
    public ListNode reverseBetween(ListNode head, int m, int n) {
        ListNode p = head;
        ListNode last = null;
        ListNode tmp = null;
        int i = 1;
        ListNode left_last = null;
        ListNode right_begin = null;
        ListNode reverseEnd = null;
        ListNode reverseBegin = null;
        while (p != null) {
            if(i==m-1) left_last = p;
            if(i==n+1) right_begin = p;
            if(i==m) reverseEnd = p;
            if(i==n) reverseBegin = p;
            if((i<=m) || (i>n)) {
                last = p;
                p = p.next;
            }
            else {
                tmp = p.next;
                p.next = last;
                last = p;
                p = tmp;
            }
            i++;
        }
        if(left_last==null) head = reverseBegin;
        else left_last.next = reverseBegin;
        reverseEnd.next = right_begin;
        return head;
    }
}