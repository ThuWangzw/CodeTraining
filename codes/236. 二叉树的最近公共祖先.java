/*
先找到最近公共祖先的特征
 */

class Solution {
    TreeNode result = null;
    public boolean target_in_tree(TreeNode root, TreeNode p, TreeNode q) {
        if(root == null) return false;
        boolean in_left = target_in_tree(root.left, p, q);
        boolean in_right = target_in_tree(root.right, p, q);
        if((in_left && in_right) || (((root == p) || (root == q)) && (in_left||in_right))){
            result = root;
        }
        return in_left || in_right || (root==p) || (root==q);
    }
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        target_in_tree(root, p, q);
        return result;
    }
}