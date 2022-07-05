class Solution {
public:
    int minimumTotal(vector<vector<int>>& triangle) {
        int n = triangle.size();
        vector<int> ans = vector<int>(n);
        ans[0] = triangle[0][0];
        for(int i=1; i<n; i++) {
            ans[i] = ans[i-1] + triangle[i][i];
            for(int j=i-1; j>0; j--) {
                ans[j] = min(ans[j-1], ans[j]) + triangle[i][j];
            }
            ans[0] += triangle[i][0];
        }
        int minpath = INT32_MAX;
        for(int i=0; i<n; i++) {
            minpath = min(minpath, ans[i]);
        }
        return minpath;
    }
};