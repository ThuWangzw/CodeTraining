#include <iostream>
#include <fstream>
#include <string>

using namespace std;
unsigned long long len[int(1e5)+1];
string f0 = "What are you doing at the end of the world? Are you busy? Will you save us?";
string a = "What are you doing while sending \"";
string b = "\"? Are you busy? Will you send \"";
string c = "\"?";

char getCharFromF(int n, unsigned long long k) {
    if(k>len[n]) return '.';
    k--;
    if(n==0) return f0[k];
    if(k<a.size()) return a[k];
    else if(k<a.size()+len[n-1]) return getCharFromF(n-1, k-a.size()+1);
    else if(k<a.size()+len[n-1]+b.size()) return b[k-(a.size()+len[n-1])];
    else if(k<a.size()+len[n-1]+b.size()+len[n-1]) return getCharFromF(n-1, k-(a.size()+len[n-1]+b.size())+1);
    else return c[k-(a.size()+len[n-1]+b.size()+len[n-1])];
}

void CF896A(istream &cin, ostream &cout) {
    
    string ans;
    len[0] = f0.size();
    for(int i=1; i<=int(1e5); i++) len[i]=len[i-1]*2+a.size()+b.size()+c.size();
    int t;
    for(cin>>t; t>0; t--) {
        int n;
        unsigned long long k;
        cin >> n >> k;
        ans.push_back(getCharFromF(n, k));
    }
    cout << ans << '\n';
}


int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF896A(cin, cout);
    return 0;
}