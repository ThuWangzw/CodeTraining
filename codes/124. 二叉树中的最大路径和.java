class Solution {
    int result = Integer.MIN_VALUE;
    int maxContribution(TreeNode root) {
        if(root == null) return 0;
        int leftContribution = maxContribution(root.left);
        int rightContribution = maxContribution(root.right);
        result = Math.max(result, root.val+Math.max(leftContribution, 0)+Math.max(rightContribution, 0));
        return Math.max(Math.max(leftContribution, 0), rightContribution)+root.val;
    }
    public int maxPathSum(TreeNode root) {
        maxContribution(root);
        return result;
    }
}