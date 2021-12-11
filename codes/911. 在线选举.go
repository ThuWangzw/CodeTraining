// 二分查找
type TopVotedCandidate struct {
	candidates []int
	times      []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	n := len(persons)
	candidates := make([]int, n)
	personCnt := make(map[int]int)
	maxperson := -1
	maxpersonCnt := 0
	for i, person := range persons {
		personCnt[person]++
		if personCnt[person] >= maxpersonCnt {
			maxperson = person
			maxpersonCnt = personCnt[person]
		}
		candidates[i] = maxperson
	}
	return TopVotedCandidate{
		candidates: candidates,
		times:      times,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	begin := 0
	end := len(this.times) - 1
	for begin < end {
		mid := (begin + end + 1) / 2
		if this.times[mid] > t {
			end = mid - 1
		} else {
			begin = mid
		}
	}
	return this.candidates[begin]
}