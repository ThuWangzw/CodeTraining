class Solution {
public:
    void nextPermutation(vector<int>& nums) {
        auto n = nums.size();
        if(n<=1) {
            return;
        }
        int begin = n-2;
        for(; begin>=0; begin--) {
            if(nums[begin] >= nums[begin+1]) {
                continue;
            }
            break;
        }
        sort(nums.begin()+begin+1, nums.end());
        if(begin==-1) {
            return;
        }
        auto pivot = nums[begin];
        int idx = upper_bound(nums.begin()+begin+1, nums.end(), pivot)-nums.begin();
        nums[begin] = nums[idx];
        nums[idx] = pivot;
    }
};