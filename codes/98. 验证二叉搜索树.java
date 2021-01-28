/*
 我采用的是递归判断是否是二叉平衡树
 一个更为优雅的方法是利用性质：二叉搜索树的中序遍历一定是不降的序列。因此进行一遍中序遍历即可。
 */
class Solution {
    static class TreeProp{
        int min_val;
        int max_val;
        boolean is_balance;
        TreeProp(int min_val, int max_val, boolean is_balance) {
            this.min_val = min_val;
            this.max_val = max_val;
            this.is_balance = is_balance;
        }
        TreeProp(boolean is_balance) {
            this.is_balance = is_balance;
        }
    }
    TreeProp isBalance(TreeNode root) {
        if(root==null) return null;
        TreeProp left_prop = isBalance(root.left);
        if((left_prop != null) && ((!left_prop.is_balance) || (root.val<=left_prop.max_val))) {
            return new TreeProp(false);
        }
        TreeProp right_prop = isBalance(root.right);
        if((right_prop != null) && ((!right_prop.is_balance) || (root.val>=right_prop.min_val))) {
            return new TreeProp(false);
        }
        return new TreeProp(left_prop == null ? root.val : left_prop.min_val, right_prop == null ? root.val : right_prop.max_val, true);
    }
    public boolean isValidBST(TreeNode root) {
        return isBalance(root).is_balance;
    }
}