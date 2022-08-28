class Solution {
public:
    vector<vector<int>> buildMatrix(int k, vector<vector<int>>& rowConditions, vector<vector<int>>& colConditions) {
        auto rows = typoSort(k, rowConditions);
        auto cols = typoSort(k, colConditions);
        if(rows.size()==0 || cols.size()==0) return vector<vector<int>>();
        vector<vector<int>> mat(k);
        for(int i=0; i<k; i++) {
            mat[i] = vector<int>(k, 0);
        }
        for(int i=0; i<k; i++) {
            mat[rows[i]][cols[i]] = i+1;
        }
        return mat;
    }

    vector<int> typoSort(int k, vector<vector<int>>& conditions) {
        vector<vector<int>> graph(k);
        vector<int> cnts(k);
        for(auto cond : conditions) {
            graph[cond[0]-1].push_back(cond[1]-1);
            cnts[cond[1]-1]++;
        }
        queue<int> qu;
        for(int i=0; i<k; i++) {
            if(cnts[i]==0) qu.push(i);
        }
        int i=0;
        vector<int> ans(k);
        while(qu.size()>0) {
            auto front = qu.front();
            qu.pop();
            ans[front] = i;
            i++;
            for(auto child : graph[front]) {
                cnts[child]--;
                if(cnts[child]==0) {
                    qu.push(child);
                }
            }
        }
        if(i!=k) return vector<int>();
        return ans;
    }
};