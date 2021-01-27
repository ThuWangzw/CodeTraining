class Solution {
    public int height_or_unbalance(TreeNode root) {
        if(root == null) return 0;
        int left_height = height_or_unbalance(root.left);
        if(left_height==-1) return -1;
        int right_height = height_or_unbalance(root.right);
        if((right_height==-1) || (Math.abs(left_height-right_height)>1)) {
            return -1;
        }
        return Math.max(left_height, right_height)+1;
    }
    public boolean isBalanced(TreeNode root) {
        return height_or_unbalance(root)!=-1;
    }
}