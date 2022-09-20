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

void CF1395C(istream &cin, ostream &cout) {
    int n, m;
    cin >> n >> m;
    vector<int> aa(n);
    vector<vector<int>> remains(n, vector<int>(m));
    for(int i=0; i<n; i++) cin >> aa[i];
    for(int i=0; i<m; i++) {
       int b;
       cin >> b;
       for(int j=0; j<n; j++) remains[j][i] = b;
    }
    int ans = 0;
    for(int bit=10; bit>=0; bit--) {
        bool isone = false;
        for(int i=0; i<n; i++) {
            int a=aa[i];
            if(((aa[i]>>bit)&1)==0) continue;
            bool haszero = false;
            for(auto b : remains[i]) {
                if(((b>>bit)&1)==0) {
                    haszero = true;
                    break;
                }
            }
            if(!haszero) {
                isone = true;
                break;
            }
        }
        if(isone) {
            ans |= 1<<bit;
            continue;
        }
        for(int i=0; i<n; i++) {
            if(((aa[i]>>bit)&1)==0) continue;
            auto it = remains[i].begin();
            while(it != remains[i].end()) {
                if((((*it)>>bit)&1)==1) {
                    it = remains[i].erase(it);
                } else {
                    it++;
                }
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
    CF1395C(cin, cout);
    return 0;
}