class Solution {
public:
    int garbageCollection(vector<string>& garbage, vector<int>& travel) {
        vector<int> remainTime(3);
        vector<int> alltime(3);
        remainTime[0]=remainTime[1]=remainTime[2] = 0;
        for(int i=0; i<garbage.size(); i++) {
            vector<int> cnts(3);
            cnts[0]=cnts[1]=cnts[2]=0;
            for(auto c : garbage[i]) {
                if(c=='M') {
                    cnts[0]++;
                } else if(c=='P') {
                    cnts[1]++;
                } else {
                    cnts[2]++;
                }
            }
            for(int j=0; j<3; j++) {
                if(i>0) remainTime[j] += travel[i-1];
                if(i==0 || cnts[j]>0) {
                    alltime[j] += remainTime[j]+cnts[j];
                    remainTime[j] = 0;
                }
            }
        }
        return alltime[0]+alltime[1]+alltime[2];
    }
};