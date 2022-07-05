class Solution {
public:
    int minPathSum(vector<vector<int>>& grid) {
        auto m = grid.size();
        auto n = grid[0].size();
        vector<int> minpath = vector<int>(n);
        for(int i=0; i<n; i++) {
            if(i==0) {
                minpath[0] = grid[0][0];
            } else {
                minpath[i] = minpath[i-1]+grid[0][i];
            }
        }
        for(int i=1; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(j==0) {
                    minpath[0] += grid[i][0];
                } else {
                    minpath[j] = min(minpath[j], minpath[j-1]) + grid[i][j];
                }
            }
        }
        return minpath[n-1];
    }
};