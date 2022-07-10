class SmallestInfiniteSet {
private:
    priority_queue<int, vector<int>, greater<int>> nums;
    unordered_set<int> numMap;
public:
    SmallestInfiniteSet() {
        for(int i=1; i<=1001; i++) {
            nums.push(i);
            numMap.insert(i);
        }
    }
    
    int popSmallest() {
        int top = nums.top();
        nums.pop();
        numMap.erase(top);
        return top;
    }
    
    void addBack(int num) {
        if(numMap.find(num) != numMap.end()) {
            return;
        }
        nums.push(num);
        numMap.insert(num);
    }
};

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * SmallestInfiniteSet* obj = new SmallestInfiniteSet();
 * int param_1 = obj->popSmallest();
 * obj->addBack(num);
 */