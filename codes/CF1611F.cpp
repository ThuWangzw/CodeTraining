#include <cstdio>
#include <vector>
#include <iostream>
#include <fstream>

using namespace std;

void CF1611F(istream& cin, ostream& cout) {
	int t;
	for(cin >> t; t>0; t--) {
		int n, s;
		cin >> n >> s;
		vector<int> nums(n);
		long long sum = 0;
		int left=0;
		int l=-1, r=-2;
		for(int i=0; i<n; i++) {
			cin >> nums[i];
			sum += nums[i];
			while(sum+s<0 && left<=i) {
				sum -= nums[left];
				left++;
			}

			if (sum+s >= 0 && r-l < i-left) {
				l = left+1;
				r = i+1;
			}
		}
		if(l==-1) {
			cout << l << endl;
		} else {
			cout << l << ' ' << r << endl;
		}
	}
	
}

int main() {
	std::ios::sync_with_stdio(false);
	std::cin.tie(nullptr);
	// ifstream infile("src/input.txt");
	// cin.rdbuf(infile.rdbuf());

    CF1611F(cin, cout);
	return 0;
}