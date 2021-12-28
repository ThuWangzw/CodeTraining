// Trie树+深搜
type byLen []string

func (bylen byLen) Len() int {
	return len(bylen)
}

func (bylen byLen) Swap(i, j int) {
	bylen[i], bylen[j] = bylen[j], bylen[i]
}

func (bylen byLen) Less(i, j int) bool {
	return len(bylen[i]) < len(bylen[j])
}

type Trie struct {
	children [26]*Trie
	isend    bool
}

func (trie *Trie) insert(s string) {
	if len(s) == 0 {
		trie.isend = true
	} else {
		if trie.children[s[0]-'a'] == nil {
			trie.children[s[0]-'a'] = &Trie{}
		}
		trie.children[s[0]-'a'].insert(s[1:])
	}
}

func isConcatWord(word string, trie *Trie) bool {
	n := len(word)
	if n == 0 {
		return trie.isend
	}
	node := trie
	for i := 0; i < n; i++ {
		if i > 0 && node.isend {
			flag := isConcatWord(word[i:], trie)
			if flag {
				return true
			}
		}
		if node.children[word[i]-'a'] != nil {
			node = node.children[word[i]-'a']
		} else {
			return false
		}
	}
	return node.isend
}

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Sort(byLen(words))
	ans := make([]string, 0)
	trie := &Trie{}
	for _, word := range words {
		if isConcatWord(word, trie) {
			ans = append(ans, word)
		} else {
			trie.insert(word)
		}
	}
	return ans
}