#include <iostream>
#include <fstream>
#include <string>
#include <queue>
#include <vector>
#include <algorithm>
using namespace std;

void getBits(long long n, vector<int> &nums) {
    while(n>0) {
        nums.push_back(n&1);
        n = n>>1;
    }
}

void CF484A(istream &cin, ostream &cout) {
    int t;
    long long a  = 1;
    for(cin>>t; t>0; t--) {
        long long left, right;
        cin >> left >> right;
        vector<int> lefts;
        vector<int> rights;
        getBits(left, lefts);
        getBits(right, rights);
        if(right==(a<<rights.size())-1) {
            cout << right << endl;
            continue;
        }
        if(rights.size()>lefts.size()) {
            cout << (a<<(rights.size()-1))-1 << endl;
            continue;
        }
        long long ans = 0;
        for(long long i=rights.size()-1; i>=0; i--) {
            if(lefts[i]==1) {
                ans = ans*2+1;
            } else if(rights[i]==0) {
                ans *= 2;
            } else {
                ans = (ans<<(i+1)) | ((a<<(i+1))-1);
                if(ans>right) {
                    ans -= a<<i;
                }        
                break;      
            }
        }
        cout << ans << endl;
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF484A(cin, cout);
}