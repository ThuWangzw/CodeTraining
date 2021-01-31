/*
    方法一：使用哈希表记录节点是否被访问过，时间复杂度On，空间复杂度On
 */

class HashSolution {
    public boolean hasCycle(ListNode head) {
        HashMap<Integer, ListNode> node_hash = new HashMap<>();
        ListNode p = head;
        boolean has_cycle = false;
        while (p != null) {
            Integer p_hash = p.hashCode();
            if(node_hash.containsKey(p_hash)) {
                has_cycle = true;
                break;
            }
            node_hash.put(p_hash, p);
            p = p.next;
        }
        return has_cycle;
    }
}

/*
    方法2： 使用Floyd判圈算法（龟兔赛跑算法），该算法不仅可以判圈，而且可以求出圈的长度和起点，时间复杂度On，空间复杂度O1
 */
public class FloydSolution {
    public boolean hasCycle(ListNode head) {
        if(head==null) return false;
        ListNode fast = head;
        ListNode slow = head;
        while (fast.next != null && fast.next.next!=null) {
            fast = fast.next.next;
            slow = slow.next;
            if(fast==slow) return true;
        }
        return false;
    }
}