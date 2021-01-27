/*
 二叉树的层次遍历，非递归形式(BFS)
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> nodes = new ArrayList<>();
        LinkedList<TreeNode> queue = new LinkedList<>();
        if(root==null) return nodes;
        queue.add(root);
        while (!queue.isEmpty()) {
            int queue_len = queue.size();
            List<Integer> nodes_in_this_level = new ArrayList<Integer>();
            for(int i=0; i<queue_len; i++) {
                TreeNode node = queue.pop();
                nodes_in_this_level.add(node.val);
                if(node.left != null) queue.add(node.left);
                if(node.right != null) queue.add(node.right);
            }
            nodes.add(nodes_in_this_level);
        }
        return nodes;
    }
}