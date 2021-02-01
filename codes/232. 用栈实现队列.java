class MyQueue {
    ArrayList<Integer> s1;
    ArrayList<Integer> s2;
    Integer front;

    void s1tos2() {
        while (!s1.isEmpty()) {
            s2.add(s1.remove(s1.size()-1));
        }
    }

    void s2tos1() {
        while (!s2.isEmpty()) {
            s1.add(s2.remove(s2.size()-1));
        }
    }

    /** Initialize your data structure here. */
    public MyQueue() {
        s1 = new ArrayList<>();
        s2 = new ArrayList<>();
    }

    /** Push element x to the back of queue. */
    public void push(int x) {
        if(!s2.isEmpty()) s2tos1();
        s1.add(x);
        if(s1.size() == 1) front = s1.get(0);
    }

    /** Removes the element from in front of queue and returns that element. */
    public int pop() {
        if(s2.isEmpty()) s1tos2();
        Integer pop_value = s2.remove(s2.size()-1);
        if(s2.size() > 0) front = s2.get(s2.size()-1);
        else front = null;
        return pop_value;
    }

    /** Get the front element. */
    public int peek() {
        return front;
    }

    /** Returns whether the queue is empty. */
    public boolean empty() {
        return s1.isEmpty() && s2.isEmpty();
    }
}