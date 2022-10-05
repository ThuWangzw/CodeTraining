#include <iostream>
#include <fstream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <algorithm>
using namespace std;
 
void CF811C(istream &cin, ostream &cout) {
    int n;
    cin>>n;
    vector<int> nums(n);
    unordered_map<int,pair<int, int>> poses;
    vector<vector<bool>> ia(n);
    for(int i=0; i<n; i++) {
        cin >> nums[i];
        if(poses.find(nums[i]) == poses.end()) {
            poses[nums[i]] = make_pair(i, i);
        }
        poses[nums[i]].second = i;
        ia[i] = vector<bool>(n, false);
    }
    for(int i=0; i<n; i++) {
        int j=i;
        int aj = i;
        while(j<n) {
            if(poses[nums[j]].first < i) {
                break;
            }
            if(poses[nums[j]].second > j) {
                aj = max(aj, poses[nums[j]].second);
            }
            if(j>=aj) ia[i][j] = true;
            j++;
        }
    }
    
    vector<int> dp(n, 0);
    int ans = 0;
    for(int i=0; i<n; i++) {
        unordered_set<int> oc;
        int res = 0;
        if(i>0) dp[i]=dp[i-1];
        for(int j=i; j>=0; j--) {
            if(oc.find(nums[j]) == oc.end()) {
                oc.insert(nums[j]);
                res ^= nums[j];
            }
            if(!ia[j][i]) continue;
            if(j>=1) {
                dp[i] = max(dp[i], dp[j-1]+res);
            } else {
                dp[i] = max(dp[i], res);
            }
        }
        ans = max(ans, dp[i]);
    }
    cout << ans << '\n';
}
 
int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);
 
    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF811C(cin, cout);
}