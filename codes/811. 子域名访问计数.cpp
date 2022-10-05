class Solution {
public:
    vector<string> subdomainVisits(vector<string>& cpdomains) {
        unordered_map<string,int> cnts;
        for(auto rawstr : cpdomains) {
            auto sp = rawstr.find(' ');
            auto cnt = atoi(rawstr.substr(0, sp).c_str());
            auto lastidx = sp;
            while(lastidx != string::npos) {
                auto domain = rawstr.substr(lastidx+1);
                if(cnts.find(domain) == cnts.end()) {
                    cnts[domain] = 0;
                }
                cnts[domain] += cnt;
                lastidx = rawstr.find('.', lastidx+1);
            }
        }
        vector<string> ans;
        for(auto item : cnts) {
            ans.push_back(to_string(item.second)+" "+item.first);
        }
        return ans;
    }
};