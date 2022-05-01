type Queue []byte

func QueueConstructor() Queue {
	return make([]byte, 0)
}

func (queue *Queue) Push(c byte) {
	(*queue) = append((*queue), c)
}

func (queue *Queue) Pop() {
	(*queue) = (*queue)[1:]
}

func (queue *Queue) Top() byte {
	return (*queue)[0]
}

func (queue *Queue) Len() int {
	return len(*queue)
}

func firstUniqChar(s string) byte {
	cMap := make(map[byte]int)
	queue := QueueConstructor()
	for i := 0; i < len(s); i++ {
		if cMap[s[i]] > 0 {
			cMap[s[i]]++
			for queue.Len() > 0 {
				top := queue.Top()
				if cMap[top] > 1 {
					queue.Pop()
				} else {
					break
				}
			}
		} else {
			cMap[s[i]] = 1
			queue.Push(s[i])
		}
	}
	if queue.Len() == 0 {
		return ' '
	}
	return queue.Top()
}