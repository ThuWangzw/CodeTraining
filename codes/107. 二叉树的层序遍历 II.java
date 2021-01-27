/*
 把层次遍历逆过来就好了。此题还有BFS的解法，每次遍历到某个节点时，就把该节点放到结果数组中相应的位置即可。
 */
class Solution {
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
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
            nodes.add(0, nodes_in_this_level);
        }
        return nodes;
    }
}