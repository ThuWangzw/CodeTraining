#include <cstdio>
#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

void CF933A(istream &cin, ostream &cout) {
	int n;
	int dp[4] = {0,0,0,0};
	for(cin>>n; n>0; n--) {
		int num;
		cin >> num;
		if(num==1) {
			dp[0]++;
			dp[2] = max(dp[2]+1, dp[1]+1);
		} else {
			dp[1] = max(dp[1]+1, dp[0]+1);
			dp[3] = max(dp[3]+1, dp[2]+1);
		}
	}
	cout << max(max(dp[0], dp[1]), max(dp[2], dp[3])) << endl;
}

int main() {
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	// ifstream ifile("src/input.txt");
	// cin.rdbuf(ifile.rdbuf());
	CF933A(cin, cout);
}