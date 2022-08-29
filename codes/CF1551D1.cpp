#include <iostream>
#include <fstream>
using namespace std;

#define FLAG(flag) ((flag)?"YES":"NO")

void CF1551D1(istream &cin, ostream &cout) {
    int t;
    for(cin>>t; t>0; t--) {
        int n,m,k;
        cin >>n>>m>>k;
        if(k>(n*m/2)) {
            cout << FLAG(false) << endl;
        }
        if(n%2!=0) {
            cout << FLAG((m%2==0 && (k-m/2)>=0 && (k-m/2)%2==0)) << endl;
        } else {
            cout << FLAG((n%2==0 && k%2==0 && (k<=(m/2*n)))) <<endl;
        }
    }
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF1551D1(cin, cout);
}