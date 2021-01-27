class Solution {
    // 递归解法
    public int maxDepth(TreeNode root) {
        if(root == null){
            return 0;
        }
        return Math.max(maxDepth(root.left), maxDepth(root.right))+1;
    }

    // BFS解法
    public int maxDepth_1(TreeNode root) {
        if(root == null) {
            return 0;
        }
        LinkedList<TreeNode> queue = new LinkedList<>();
        int level = 0;
        queue.add(root);
        while (!queue.isEmpty()) {
            level++;
            int queue_len = queue.size();
            for (int i=0; i<queue_len; i++){
                TreeNode node = queue.pop();
                if(node.left != null) queue.add(node.left);
                if(node.right != null) queue.add(node.right);
            }
        }
        return level;
    }
}