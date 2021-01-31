class Solution {
    public int evalRPN(String[] tokens) {
        ArrayList<Integer> stack = new ArrayList<>();
        Integer right = null;
        Integer left = null;
        for (String token : tokens) {
            switch (token) {
                case "+" -> {
                    right = stack.remove(stack.size() - 1);
                    left = stack.remove(stack.size() - 1);
                    stack.add(left + right);
                }
                case "-" -> {
                    right = stack.remove(stack.size() - 1);
                    left = stack.remove(stack.size() - 1);
                    stack.add(left - right);
                }
                case "*" -> {
                    right = stack.remove(stack.size() - 1);
                    left = stack.remove(stack.size() - 1);
                    stack.add(left * right);
                }
                case "/" -> {
                    right = stack.remove(stack.size() - 1);
                    left = stack.remove(stack.size() - 1);
                    stack.add(left / right);
                }
                default -> stack.add(Integer.valueOf(token));
            }
        }
        assert stack.size()==1;
        return stack.get(0);
    }
}