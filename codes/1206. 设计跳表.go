var MAXLEVEL = 32

type Node struct {
	Val     int
	Forward []*Node
}

type Skiplist struct {
	Head  *Node
	Level int
}

func randomLevel() int {
	lv := 0
	for lv < MAXLEVEL-1 && rand.Float64() < 0.5 {
		lv++
	}
	return lv
}

func Constructor() Skiplist {
	return Skiplist{
		Head:  &Node{Val: -1, Forward: make([]*Node, MAXLEVEL)},
		Level: 0,
	}
}

func (this *Skiplist) Search(target int) bool {
	cur := this.Head
	for level := this.Level; level >= 0; level-- {
		for cur.Forward[level] != nil && cur.Forward[level].Val < target {
			cur = cur.Forward[level]
		}
	}
	cur = cur.Forward[0]
	return cur != nil && cur.Val == target
}

func (this *Skiplist) Add(num int) {
	updates := make([]*Node, MAXLEVEL)
	for i := 0; i < MAXLEVEL; i++ {
		updates[i] = this.Head
	}

	cur := this.Head
	for level := this.Level; level >= 0; level-- {
		for cur.Forward[level] != nil && cur.Forward[level].Val < num {
			cur = cur.Forward[level]
		}
		updates[level] = cur
	}

	lv := randomLevel()
	if lv > this.Level {
		this.Level = lv
	}
	node := &Node{
		Val:     num,
		Forward: make([]*Node, lv+1),
	}
	for lv >= 0 {
		node.Forward[lv] = updates[lv].Forward[lv]
		updates[lv].Forward[lv] = node
		lv--
	}
}

func (this *Skiplist) Erase(num int) bool {
	updates := make([]*Node, MAXLEVEL)
	for i := 0; i < MAXLEVEL; i++ {
		updates[i] = this.Head
	}

	cur := this.Head
	for level := this.Level; level >= 0; level-- {
		for cur.Forward[level] != nil && cur.Forward[level].Val < num {
			cur = cur.Forward[level]
		}
		updates[level] = cur
	}
	if cur.Forward[0] == nil || cur.Forward[0].Val != num {
		return false
	}
	lv := this.Level
	for lv >= 0 {
		if updates[lv].Forward[lv] != nil && updates[lv].Forward[lv].Val == num {
			updates[lv].Forward[lv] = updates[lv].Forward[lv].Forward[lv]
		}
		lv--
	}
	return true
}