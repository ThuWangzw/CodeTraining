/*
 floyd判圈算法（龟兔赛跑算法）
 */
public class Solution {
    public ListNode detectCycle(ListNode head) {
        if(head==null) return null;
        // 1. has cycle?
        boolean has_cycle = false;
        ListNode fast = head;
        ListNode slow = head;
        while (fast.next!=null && fast.next.next != null) {
            fast = fast.next.next;
            slow = slow.next;
            if(slow==fast) {
                has_cycle = true;
                break;
            }
        }
        if(!has_cycle) return null;
        // 2. find cycle's starting point
        slow = head;
        while (slow != fast) {
            slow = slow.next;
            fast = fast.next;
        }
        return slow;
    }
}