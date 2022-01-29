// 一个很妙的trick是数组里面永远删除最后一个
type RandomizedSet struct {
	elements []int
	elementMap map[int]int
}


func Constructor() RandomizedSet {
	return RandomizedSet{
		elements: make([]int, 0),
		elementMap: make(map[int]int),
	}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _,in := this.elementMap[val]; in {
		return false
	}
	this.elementMap[val] = len(this.elements)
	this.elements = append(this.elements, val)
	return true
}


func (this *RandomizedSet) Remove(val int) bool {
	if idx,in := this.elementMap[val]; in {
		endval := this.elements[len(this.elements)-1]
		endidx := len(this.elements)-1
		this.elementMap[endval], this.elementMap[val] = this.elementMap[val], this.elementMap[endval]
		this.elements[idx], this.elements[endidx] = this.elements[endidx], this.elements[idx]
		delete(this.elementMap, val)
		this.elements = this.elements[:len(this.elements)-1]
		return true
	}
	return false
}


func (this *RandomizedSet) GetRandom() int {
	n := len(this.elements)
	idx := rand.Intn(n)
	return this.elements[idx]
}