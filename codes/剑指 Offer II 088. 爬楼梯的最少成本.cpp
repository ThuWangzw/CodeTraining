class Solution {
	public:
		int minCostClimbingStairs(vector<int>& cost) {
			auto n = cost.size();
			if(n<=1) {
				return 0;
			}
			vector<int> sum(n+1, 0);
			for(int i=2; i<=n; i++) {
				sum[i] = min(sum[i-1]+cost[i-1], sum[i-2]+cost[i-2]);
			}
			return sum[n];
		}
	};