class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> nodes = new ArrayList<>();
        LinkedList<TreeNode> queue = new LinkedList<>();
        if(root==null) return nodes;
        queue.add(root);
        int level = 0;
        while (!queue.isEmpty()) {
            int queue_len = queue.size();
            List<Integer> nodes_in_this_level = new ArrayList<Integer>();
            LinkedList<TreeNode> tmp_queue = new LinkedList<>();
            if(level%2==0){
                for(int i=0; i<queue_len; i++) {
                    TreeNode node = queue.pop();
                    nodes_in_this_level.add(node.val);
                    if(node.left != null) tmp_queue.add(0, node.left);
                    if(node.right != null) tmp_queue.add(0, node.right);
                }
            }
            else {
                for(int i=0; i<queue_len; i++) {
                    TreeNode node = queue.pop();
                    nodes_in_this_level.add(node.val);
                    if(node.right != null) tmp_queue.add(0, node.right);
                    if(node.left != null) tmp_queue.add(0, node.left);
                }
            }
            queue = tmp_queue;
            nodes.add(nodes_in_this_level);
            level++;
        }
        return nodes;
    }
}