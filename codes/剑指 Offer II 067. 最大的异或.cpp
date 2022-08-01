class Trie {
    vector<Trie*> node;
    bool end;
public:
    /** Initialize your data structure here. */
    Trie() {
        node = vector<Trie*>(2);
        for(int i=0; i<2; i++) node[i] = nullptr;
        end = false;
    }
    
    /** Inserts a word into the trie. */
    void insert(int num, int i) {
        if(i == -1) {
            end = true;
            return;
        }
        int idx = (num>>i) & 1;
        if(node[idx]==nullptr) {
            node[idx] = new Trie();
        }
        node[idx]->insert(num, i-1);
    }
    
    /** Returns if the word is in the trie. */
    int search(int num, int i) {
        if(i==-1) {
            return 0;
        }
        int flag = (num>>i) & 1;
        int target = 1-flag;
        if(node[target]==nullptr) {
            return node[flag]->search(num, i-1);
        } else {
            return 1<<i | node[target]->search(num, i-1);
        }
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */

class Solution {
public:
    int findMaximumXOR(vector<int>& nums) {
        auto trie = new Trie();
        for(int num : nums) {
            trie->insert(num, 31);
        }
        int maxans = 0;
        for(auto num : nums) {
            maxans = max(maxans, trie->search(num, 31));
        }
        return maxans;
    }
};