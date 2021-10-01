// 蓄水池抽样算法
type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	i := 1
	var sp *ListNode
	p := this.head
	for p != nil {
		if rand.Float32() < (float32(1) / float32(i)) {
			sp = p
		}
		i++
		p = p.Next
	}
	return sp.Val
}