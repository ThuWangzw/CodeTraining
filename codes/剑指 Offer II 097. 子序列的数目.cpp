class Solution {
public:
    int numDistinct(string s, string t) {
        int m = s.size();
        int n = t.size();
        vector<vector<long long>> mat(m+1);
        for(int i=0; i<=m; i++) {
            mat[i] = vector<long long>(n+1);
            mat[i][0] = 1;
        }
        for(int i=1; i<=m; i++) {
            for(int j=max(1, i+n-m); j<=n; j++) {
                mat[i][j] = mat[i-1][j];
                if(s[i-1]==t[j-1]) {
                    mat[i][j] += mat[i-1][j-1];
                }
            }
        }
        return mat[m][n];
    }
};