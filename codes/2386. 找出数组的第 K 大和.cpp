struct Cmp{
    bool operator()(pair<long long, int> &a, pair<long long, int> &b) {
        return a.first < b.first;
    }
};

class Solution {
public:
    long long kSum(vector<int>& nums, int k) {
        long long sum = 0;
        for(int i=0; i<nums.size(); i++) {
            if(nums[i]>0) sum += nums[i];
            else nums[i] = -nums[i];
        }
        sort(nums.begin(), nums.end());
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, Cmp> pq;
        pq.push(make_pair(sum, 0));
        for(int i=1; i<k; i++) {
            auto top = pq.top();
            pq.pop();
            if(top.second < nums.size()) {
                pq.push(make_pair(top.first-nums[top.second], top.second+1)) ;
                if(top.second>0) {
                    pq.push(make_pair(top.first-nums[top.second]+nums[top.second-1], top.second+1));
                }
            }
        }
        return pq.top().first;
    }
};