class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        ArrayList<TreeNode> stack = new ArrayList<>();
        ArrayList<Integer> travel_result = new ArrayList<>();
        TreeNode p = root;
        while ((stack.size()>0) || (p != null)) {
            if(p != null) {
                stack.add(p);
                p = p.left;
            }
            else {
                p = stack.remove(stack.size()-1);
                travel_result.add(p.val);
                p = p.right;
            }
        }
        return travel_result;
    }
}