class Solution {
public:
    int maximalRectangle(vector<string>& matrix) {
        int m = matrix.size();
        if(m==0) return 0;
        int n = matrix[0].size();
        int maxans = 0;
        vector<int> row(n);
        vector<int> area(n);
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(matrix[i][j]=='0') row[j]=0;
                else row[j]++;
            }
            vector<int> st;
            for(int j=0; j<n; j++) {
                while(st.size()>0 && row[st.back()]>=row[j]) {
                    st.pop_back();
                }
                if(st.size()>0) {
                    area[j] = (j-st.back())*row[j];
                } else {
                    area[j] = (j+1)*row[j];
                }
                st.push_back(j);
            }
            st.clear();
            for(int j=n-1; j>=0; j--) {
                while(st.size()>0 && row[st.back()]>=row[j]) {
                    st.pop_back();
                }
                if(st.size()>0) {
                    area[j] += (st.back()-j-1)*row[j];
                } else {
                    area[j] += (n-j-1)*row[j];
                }
                maxans = max(maxans, area[j]);
                st.push_back(j);
            }
        }
        return maxans;
    }
};