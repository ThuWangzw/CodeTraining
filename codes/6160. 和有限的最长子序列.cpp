class Solution {
public:
    vector<int> answerQueries(vector<int>& nums, vector<int>& queries) {
        sort(nums.begin(), nums.end());
        for(int i=0; i<nums.size(); i++) {
            if(i>0) {
                nums[i] += nums[i-1];
            }
        }
        vector<int> ans;
        for(auto q : queries) {
            auto it = upper_bound(nums.begin(), nums.end(), q);
            ans.push_back(it-nums.begin());
        }
        return ans;
    }
};