#include <iostream>
#include <fstream>
#include <vector>
#include <utility>
#include <algorithm>

using namespace std;

int CountPrime(unsigned long long n, unsigned long long p) {
    int cnt = 0;
    while(n%p==0) {
        n /= p;
        cnt++;
    }
    return cnt;
}

void CF837D(istream &cin, ostream &cout) {
    int n, k;
    cin >> n >> k;
    vector<pair<int, int>> nums(n);
    int ans = 0;
    int maxp = 26*n;
    vector<vector<int>> dp(k);
    int cnt5 = 0;
    for(int i=0; i<k; i++) dp[i] = vector<int>(maxp, -1);
    for(int i=0; i<n; i++) {
        unsigned long long num;
        cin >> num;
        auto newpair = make_pair(CountPrime(num, 2), CountPrime(num, 5));
        cnt5 += newpair.second;
        for(int j=min(i,k-1); j>=0; j--) {
            if(j==0) {
                dp[j][newpair.second] = max(dp[j][newpair.second], newpair.first);
                ans = max(ans, min(newpair.second, dp[j][newpair.second]));
                continue;
            }
            for(int p=cnt5; p>=newpair.second; p--) {
                if(dp[j-1][p-newpair.second]==-1) continue;
                auto last = newpair.first + dp[j-1][p-newpair.second];
                dp[j][p] = max(dp[j][p], last);
                ans = max(ans, min(p, dp[j][p]));
            }
        }
    }
    cout << ans << endl;
}


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/inputs.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF837D(cin, cout);
    return 0;
}