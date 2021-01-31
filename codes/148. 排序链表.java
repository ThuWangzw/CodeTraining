/*
 方法一：快速排序。因为快排不稳定，所以该方法对某些测试样例会超时
 */
class QuickSortSolution {
    static class NListNode {
        ListNode head;
        ListNode tail;
        NListNode(ListNode head, ListNode tail){
            this.head = head;
            this.tail = tail;
        }
    }

    public NListNode nsortList(ListNode head) {
        if((head==null) || (head.next == null)) return new NListNode(head, head);
        int x = head.val;

        // partition
        ListNode left_head = new ListNode(-1, null);
        ListNode left_tail = left_head;
        ListNode right_head = new ListNode(-1, null);
        ListNode right_tail = right_head;
        ListNode p = head.next;
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
        left_tail.next = null;
        right_tail.next = null;
        NListNode left = nsortList(left_head.next);
        NListNode right = nsortList(right_head.next);
        if(left.head==null) {
            head.next = right.head;
            right.head = head;
            return right;
        }
        if(right.head==null) {
            left.tail.next = head;
            head.next = null;
            left.tail = head;
            return left;
        }
        else {
            left.tail.next = head;
            head.next = right.head;
            return new NListNode(left.head, right.tail);
        }
    }

    public ListNode sortList(ListNode head) {
        return nsortList(head).head;
    }
}

/*
  方法二：自顶向下的归并排序。归并排序是稳定的nlogn复杂度，使用自顶向下的递归实现，写法比较简单，但是有logn的空间复杂度
 */

class MergeSortTopToBottomSolution {
    ListNode mergeTwoLists(ListNode l1, ListNode l2) {
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

    public ListNode sortList(ListNode head) {
        if((head==null) || (head.next==null)) return head;
        ListNode fast_p = head;
        ListNode slow_p = head;
        ListNode slow_p_last = null;
        // find mid
        while ((fast_p!=null) && (fast_p.next!=null)) {
            slow_p_last = slow_p;
            slow_p = slow_p.next;
            fast_p = fast_p.next.next;
        }
        slow_p_last.next = null;
        ListNode left = sortList(head);
        ListNode right = sortList(slow_p);
        //merge
        return mergeTwoLists(left, right);
    }
}

/*
    方法三：自底向上的归并排序。使用迭代实现的归并排序，空间复杂度为常数。
 */

class MergeSortBottomToTopSolution {
    static class NListNode {
        ListNode head;
        ListNode tail;
        NListNode(ListNode head, ListNode tail){
            this.head = head;
            this.tail = tail;
        }
    }

    NListNode mergeTwoLists(ListNode l1, ListNode l2) {
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
        while (tail.next!=null) tail = tail.next;
        return new NListNode(dummy_head.next, tail);
    }

    public ListNode sortList(ListNode head) {
        if((head==null) || (head.next==null)) return head;
        int bin_len = 1;
        ListNode slow = head;
        ListNode fast = head;
        ListNode fast_tail = head;
        ListNode slow_tail = head;
        ListNode fast_next = null;
        ListNode slow_last = null;
        boolean merge_end = false;
        while (true) {
            slow = head;
            fast = head;
            slow_last = null;
            while (fast!=null) {
                // find slow tail
                int i=1;
                slow_tail = slow;
                while ((slow_tail!=null) && (i<bin_len)){
                    slow_tail = slow_tail.next;
                    i++;
                }
                if(slow_tail==null) {
                    // no fast
                    if(slow_last==null) merge_end = true;
                    else slow_last.next = slow;
                    break;
                }

                fast = slow_tail.next;
                slow_tail.next = null;
                // find fast tail
                fast_tail = fast;
                i=1;
                while ((fast_tail!=null) && (i<bin_len)){
                    fast_tail = fast_tail.next;
                    i++;
                }
                if(fast_tail==null) fast_next = null;
                else {
                    fast_next = fast_tail.next;
                    fast_tail.next = null;
                }
                NListNode merged = mergeTwoLists(slow, fast);
                if(slow_last==null) head = merged.head;
                else slow_last.next = merged.head;
//                merged.tail.next = fast_next;
                slow_last = merged.tail;
                slow = fast = fast_next;
            }
            if(merge_end) break;
            bin_len *= 2;
        }
        return head;
    }

}