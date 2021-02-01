/*
方法一：时间复杂度为On2的遍历，容易超时
 */
class Solution {
    public int largestRectangleArea(int[] heights) {
        int max_area = 0;
        for(int i=0; i<heights.length; i++) {
            int len = 1;
            for(int j=i-1; j>=0 && heights[j] >= heights[i]; j--) {
                len++;
            }
            for(int j=i+1; j<heights.length && heights[j] >= heights[i]; j++) {
                len++;
            }
            max_area = Math.max(max_area, len*heights[i]);
        }
        return max_area;
    }
}

/*
  方法二：使用单调栈。单调栈可以找出当前元素左侧最近的比该元素小/大的元素
 */
class Solution {
    static class Rectangle{
        int index;
        int height;
        Rectangle(int index, int height) {
            this.index = index;
            this.height = height;
        }
    }
    public int largestRectangleArea(int[] heights) {
        ArrayList<Rectangle> stack = new ArrayList<>();
        int[] left_index = new int[heights.length];
        int max_area = 0;

        // find boundary
        for(int i=0; i<heights.length; i++) {
            int height = heights[i];
            while (stack.size()>0 && stack.get(stack.size()-1).height>=height){
                Rectangle removed = stack.remove(stack.size()-1);
                max_area = Math.max(max_area, heights[removed.index] * (i-left_index[removed.index]));
            }
            if(stack.size()==0) left_index[i] = 0;
            else left_index[i] = stack.get(stack.size()-1).index+1;
            stack.add(new Rectangle(i, height));
        }
        while (stack.size()>0) {
            Rectangle removed = stack.remove(stack.size()-1);
            max_area = Math.max(max_area, heights[removed.index] * (heights.length-left_index[removed.index]));
        }
        return max_area;
    }
}