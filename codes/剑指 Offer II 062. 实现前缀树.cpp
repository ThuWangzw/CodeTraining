class Trie {
    vector<Trie*> node;
    bool end;
public:
    /** Initialize your data structure here. */
    Trie() {
        node = vector<Trie*>(26);
        for(int i=0; i<26; i++) node[i] = nullptr;
        end = false;
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        if(word.size() == 0) {
            end = true;
            return;
        }
        int idx = word[0]-'a';
        if(node[idx]==nullptr) {
            node[idx] = new Trie();
        }
        node[idx]->insert(word.substr(1));
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
        if(word.size() == 0) {
            return end;
        }
        int idx = word[0]-'a';
        if(node[idx]==nullptr) return false;
        return node[idx]->search(word.substr(1));
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        if(prefix.size()==0) return true;
        int idx = prefix[0]-'a';
        if(node[idx]==nullptr) return false;
        return node[idx]->startsWith(prefix.substr(1));
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */