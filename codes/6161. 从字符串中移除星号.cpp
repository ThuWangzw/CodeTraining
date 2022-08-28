class Solution {
public:
    string removeStars(string s) {
        string news;
        for(auto c : s) {
            if(c=='*') {
                news.pop_back();
            } else {
                news.push_back(c);
            }
        }
        return news;
    }
};