/*
    需要用哈希表保存已访问过的节点
 */
class Solution {
    HashMap<Node, Node> cloned_map = new HashMap<>();
    public Node cloneGraph(Node node) {
        if(node == null) return null;
        Node cloned_node = cloned_map.get(node);
        if(cloned_node == null) {
            Node new_node = new Node(node.val);
            cloned_map.put(node, new_node);
            for(Node neighbor : node.neighbors) {
                new_node.neighbors.add(cloneGraph(neighbor));
            }
            return new_node;
        }
        return cloned_node;
    }
}