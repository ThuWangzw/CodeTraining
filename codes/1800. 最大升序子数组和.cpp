class Solution {
public:
    int maxAscendingSum(vector<int>& nums) {
        int cur=0, maxsum=0;
        for(int i=0; i<nums.size(); i++) {
            if(i==0 || nums[i]<=nums[i-1]) {
                maxsum = max(maxsum, cur);
                cur=0;
            }
            cur += nums[i];
        }
        maxsum = max(maxsum, cur);
        return maxsum;
    }
};