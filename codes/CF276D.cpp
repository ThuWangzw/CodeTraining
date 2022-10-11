#include <iostream>
#include <fstream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <algorithm>
using namespace std;

void CF276D(istream &cin, ostream &cout) {
    int n = 63;
    unsigned long long a, b;
    cin >> a >> b;
    while(n>=0 && (a>>n)==(b>>n)) n--;
    n++;
    cout << ((unsigned long long)1<<n)-1<< '\n';
}

int main() {
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    // ifstream ifs("src/input.txt");
    // cin.rdbuf(ifs.rdbuf());
    CF276D(cin, cout);
}