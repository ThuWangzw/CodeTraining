#include <iostream>
#include <fstream>
#include <string>
using namespace std;

void CF1537E1(istream &cin, ostream &cout) {
    int n, k;
    cin >> n >> k;
    string s;
    cin >> s;
    int i;
    string cmp = s+s;
    for(i=0; i<s.size(); i++) {
        string cur(s.begin()+i, s.end());
        if((cur+s)>cmp) {
            break;
        }
    }
    string ans(s.begin(), s.begin()+i);
    while(ans.size()<k) {
        ans += ans;
    }
    while(ans.size()>k) {
        ans.pop_back();
    }
    cout << ans << endl;
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF1537E1(cin, cout);
}