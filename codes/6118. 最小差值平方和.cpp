class Solution {
public:
    long long minSumSquareDiff(vector<int>& nums1, vector<int>& nums2, int k1, int k2) {
        int n = nums1.size();
        map<int, int, greater<int>> distance;
        for(int i=0; i<n; i++) {
            distance[abs(nums1[i]-nums2[i])]++;
        }
        int m = k1+k2;
        while(m>0) {
            auto top = *distance.begin();
            long long topv = top.first;
            long long topc = top.second;
            if(topv ==0) {
                break;
            }
            distance.erase(topv);
            long long nextv = 0;
            if(distance.size() > 0) {
                nextv = (*distance.begin()).first;
            }
            long long minus = m;
            if((topv-nextv)*topc<= m) {
                minus = (topv-nextv)*topc;
            }
            long long minusone = minus/topc;
            long long half = minus%topc;
            m -= minus;
            if(topc-half > 0) distance[topv-minusone] += topc-half;
            if(half>0) distance[topv-minusone-1] += half;
        }
        long long ans = 0;
        for(auto top=distance.begin(); top !=distance.end(); top++) {
            ans += ((long long)(*top).first)*((long long)(*top).first)*(*top).second;
        }
        return ans;
    }
};