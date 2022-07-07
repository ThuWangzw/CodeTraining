class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        unordered_map<char, int> cmap;
        int ans = 0;
        int len = 0;
        for(int i=0; i<s.size(); i++) {
            char c = s[i];
            if(cmap.find(c) == cmap.end()) {
                len++;
                cmap[c] = i;
            } else {
                auto lastpos = cmap[c];
                if(i-lastpos > len) {
                    len++;
                } else {
                    len = i-lastpos;
                }
                cmap[c] = i;
            }
            ans = max(ans, len);
        }
        return ans;
    }
};