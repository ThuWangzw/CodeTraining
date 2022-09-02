#include <cstdio>
#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

bool checkX(vector<int> &nums, int m, int margin) {
    int last = -1;
    for(int i=0; i<nums.size(); i++) {
        if(i==0) {
            if(nums[i]+margin>=m) {
                last = 0;
            } else {
				last = nums[i];
			}
        } else {
            if(nums[i]+margin>=m) {
                if(nums[i]+margin-m<last) {
                    last = max(last, nums[i]);
                }
            } else {
                if(nums[i]>=last) {
                    last = nums[i];
                } else if(nums[i]+margin<last) {
                    return false;
                }
            }
        }
    }
    return true;
}

void CF1168A(istream &cin, ostream &cout) {
	int n, m;
    cin >> n >> m;
    vector<int> nums(n);
    for(int i=0; i<n; i++) cin>>nums[i];
    int begin=0, end=m-1;
    while(begin<end) {
        int mid=begin+(end-begin)/2;
        if(checkX(nums, m, mid)) {
            end=mid;
        } else {
            begin=mid+1;
        }
    }
    cout << begin << endl;
}

int main() {
	ios::sync_with_stdio(false);
	cin.tie(nullptr);

	// ifstream ifile("src/input.txt");
	// cin.rdbuf(ifile.rdbuf());
	CF1168A(cin, cout);
}