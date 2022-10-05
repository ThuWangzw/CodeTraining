#include <iostream>
#include <fstream>
#include <vector>
#include <utility>
#include <string>
#include <queue>
#include <algorithm>

using namespace std;

void CF377A(istream &cin, ostream &cout) {
    int n, m, k;
    cin >> n >> m >> k;
    vector<string> mat(n);

    for(int i=0; i<n; i++) cin >> mat[i];
    // find first "."
    int br=-1, bc=-1;
    int dotcnt=0;
    for(int i=0; i<n; i++) {
        for(int j=0; j<m; j++) {
            if(mat[i][j]=='.') {
                br = i;
                bc = j;
                dotcnt++;
            }
        }
    }
    vector<vector<bool>> flags(n, vector<bool>(m, false));
    flags[br][bc] = true;
    queue<pair<int,int>> q;
    q.push(make_pair(br, bc));
    if(k+1>dotcnt) mat[br][bc] = 'X';
    int cnt=1;
    vector<vector<int>> directions = {{0,-1},{0,1},{1,0},{-1,0}};
    while(q.size()>0) {
        auto top = q.front();
        q.pop();
        for(auto direction : directions) {
            int nr=top.first+direction[0], nc=top.second+direction[1];
            if(nr>=0 && nr<n && nc>=0 && nc<m && !flags[nr][nc] && mat[nr][nc]=='.') {
                flags[nr][nc] = true;
                cnt++;
                if(cnt+k>dotcnt) mat[nr][nc] = 'X';
                q.push(make_pair(nr,nc));
            } 
        }
    }
    for(int i=0; i<n; i++) cout << mat[i] << endl;
}


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/inputs.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF377A(cin, cout);
    return 0;
}