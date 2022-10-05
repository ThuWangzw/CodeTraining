#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <string>

using namespace std;

void CF1619E(istream &cin, ostream &cout) {
    int t;
    for(cin>>t; t>0; t--) {
        int n;
        cin >> n;
        vector<int> nums(n+1, 0);
        for(int i=0; i<n; i++) {
            int num;
            cin >> num;
            nums[num]++;
        }
        bool fail = false;
        vector<int> stack;
        long long sum=0;
        for(int i=0; i<=n; i++) {
            if(fail) {
                cout << -1 << ' ';
                continue;
            }
            if(i>0) {
                while (nums[i-1]>0){
                    nums[i-1]--;
                    stack.push_back(i-1);
                }
                if(stack.size()==0) {
                    fail = true;
                    cout << -1 << ' ';
                    continue;
                }
                int num = stack.back();
                stack.pop_back();
                sum += i-1-num;
            }
            cout << sum + nums[i] << ' ';
        }
        
        cout << endl;
    }
}


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF1619E(cin, cout);
    return 0;
}