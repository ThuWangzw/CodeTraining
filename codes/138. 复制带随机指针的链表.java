/*
 官方题解：使用新旧节点交替来代替map，只用了O1的空间复杂度，非常巧妙
 */
class Solution {
    public Node copyRandomList(Node head) {
        HashMap<Integer, Integer> node_map = new HashMap<>();
        Integer i=0;
        Node n_head = null;
        ArrayList<Node> n_nodes = new ArrayList<>();
        Node p = head;
        Node last = null;
        // copy val+next
        while (p != null) {
            node_map.put(p.hashCode(), i);
            Node q = new Node(p.val);
            if(last == null) {
                n_head = last = q;
            }
            else {
                last.next = q;
                last = q;
            }
            n_nodes.add(q);
            i++;
            p = p.next;
        }

        // copy random
        p = head;
        Node q = n_head;
        while (p != null) {
            if(p.random != null) {
                Integer p_random_index = node_map.get(p.random.hashCode());
                q.random = n_nodes.get(p_random_index);
            }

            p = p.next;
            q = q.next;
        }

        return n_head;
    }
}