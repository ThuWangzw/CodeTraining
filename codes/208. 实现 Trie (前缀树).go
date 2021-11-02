type Trie struct {
	isend bool
	tree  map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		isend: false,
		tree:  make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		this.isend = true
		return
	}
	if _, ok := this.tree[word[0]]; !ok {
		tree := Constructor()
		this.tree[word[0]] = &tree
	}
	this.tree[word[0]].Insert(word[1:])
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.isend
	}
	child, ok := this.tree[word[0]]
	if !ok {
		return false
	}
	return child.Search(word[1:])
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	child, ok := this.tree[prefix[0]]
	if !ok {
		return false
	}
	return child.StartsWith(prefix[1:])
}