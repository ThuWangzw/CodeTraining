// 字典树平平无奇
type WordDictionary struct {
	trees map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{trees: make(map[byte]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	curnode := this
	idx := 0
	for idx = 0; idx < len(word); idx++ {
		_, ok := (*curnode).trees[word[idx]]
		if !ok {
			(*curnode).trees[word[idx]] = &WordDictionary{trees: make(map[byte]*WordDictionary)}
		}
		curnode = (*curnode).trees[word[idx]]
	}
	_, ok := (*curnode).trees[0]
	if !ok {
		(*curnode).trees[0] = &WordDictionary{trees: make(map[byte]*WordDictionary)}
	}
}

func (this *WordDictionary) Search(word string) bool {
	curnode := this
	idx := 0
	for idx = 0; idx < len(word); idx++ {
		if word[idx] == '.' {
			flag := false
			for ch, tree := range curnode.trees {
				if ch != 0 {
					flag = flag || tree.Search(word[idx+1:])
				}
			}
			return flag
		}
		_, ok := (*curnode).trees[word[idx]]
		if !ok {
			return false
		}
		curnode = (*curnode).trees[word[idx]]
	}
	_, ok := (*curnode).trees[0]
	return ok
}