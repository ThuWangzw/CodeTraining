/*
    方法一：使用辅助栈记录最小值
    方法二：栈中存储与最小值的差值，非常巧妙，不过需要考虑整数的边界情况，最好使用大整数
 */
class MinStack {
    ArrayList<Integer> stack;
    ArrayList<Integer> min_stack;
    Integer len;

    /** initialize your data structure here. */
    public MinStack() {
        stack = new ArrayList<>();
        min_stack = new ArrayList<>();
        len = 0;
    }

    public void push(int x) {
        stack.add(x);
        if(len>0) min_stack.add(Math.min(x, min_stack.get(len-1)));
        else min_stack.add(x);
        len++;
    }

    public void pop() {
        assert len>0;
        stack.remove(len-1);
        min_stack.remove(len-1);
        len--;
    }

    public int top() {
        assert len>0;
        return stack.get(len-1);
    }

    public int getMin() {
        return min_stack.get(len-1);
    }
}