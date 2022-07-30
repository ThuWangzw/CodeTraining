type Solution struct {
	Probs []int
	Sum   int
}

func Constructor(w []int) Solution {
	probs := make([]int, 0, len(w))
	sum := 0
	for _, a := range w {
		sum += a
		probs = append(probs, sum)
	}
	return Solution{
		Probs: probs,
		Sum:   sum,
	}
}

func (this *Solution) bs(num int) int {
	probs := this.Probs
	left, right := 0, len(probs)-1
	for left < right-3 {
		mid := (right-left)/2 + left
		if probs[mid] <= num {
			left = mid
		} else {
			right = mid
		}
	}
	for left <= right {
		if probs[left] > num {
			return left
		}
		left++
	}
	return -1
}

func (this *Solution) PickIndex() int {
	ri := rand.Intn(this.Sum)
	return this.bs(ri)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */