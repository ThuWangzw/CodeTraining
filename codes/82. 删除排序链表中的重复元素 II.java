/*
    本题可以使用dummy node方法。dummy node指的是，当链表的头部会被删除时，在链表头部插入一个不会被删除的dummy node，使得代码更加简洁
 */
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