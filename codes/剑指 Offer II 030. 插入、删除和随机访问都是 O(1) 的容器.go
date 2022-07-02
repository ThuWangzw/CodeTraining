type RandomizedSet struct {
	numIndexMap map[int]int
	nums        []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		numIndexMap: make(map[int]int),
		nums:        make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numIndexMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.numIndexMap[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	removedIdx, ok := this.numIndexMap[val]
	if !ok {
		return false
	}
	lastNum := this.nums[len(this.nums)-1]
	this.nums[removedIdx], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[removedIdx]
	this.numIndexMap[lastNum] = removedIdx
	delete(this.numIndexMap, val)
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	idx := rand.Intn(len(this.nums))
	return this.nums[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */