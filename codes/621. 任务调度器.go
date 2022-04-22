type Task struct {
	remain  int
	release int
}

type Heap struct {
	tasks   []*Task
	reverse bool
}

func (h *Heap) Push(task *Task) {
	h.tasks = append(h.tasks, task)
	h.up()
}

func (h *Heap) Pop() *Task {
	top := h.tasks[0]
	n := h.Size()
	h.tasks[0], h.tasks[n-1] = h.tasks[n-1], h.tasks[0]
	h.tasks = h.tasks[:n-1]
	h.heapify(0)
	return top
}

func (h *Heap) Top() *Task {
	return h.tasks[0]
}

func (h *Heap) Size() int {
	return len(h.tasks)
}

func (h *Heap) less(i, j int) bool {
	if h.reverse {
		// max-top
		return h.tasks[i].remain > h.tasks[j].remain
	} else {
		return h.tasks[i].release < h.tasks[j].release
	}
}

func (h *Heap) up() {
	node := h.Size() - 1
	for node > 0 {
		parent := (node - 1) / 2
		if h.less(node, parent) {
			h.tasks[node], h.tasks[parent] = h.tasks[parent], h.tasks[node]
			node = parent
		} else {
			break
		}
	}
}

func (h *Heap) heapify(nodeid int) {
	n := h.Size()
	left := nodeid*2 + 1
	right := nodeid*2 + 2
	lessNode := nodeid
	if left < n && h.less(left, lessNode) {
		lessNode = left
	}
	if right < n && h.less(right, lessNode) {
		lessNode = right
	}
	if lessNode != nodeid {
		h.tasks[nodeid], h.tasks[lessNode] = h.tasks[lessNode], h.tasks[nodeid]
		h.heapify(lessNode)
	}
}

func Constructor(tasks []*Task, reverse bool) *Heap {
	heap := &Heap{
		tasks:   tasks,
		reverse: reverse,
	}
	for i := heap.Size() - 1; i >= 0; i-- {
		heap.heapify(i)
	}
	return heap
}

func leastInterval(tasks []byte, n int) int {
	taskMap := make(map[byte]int)
	for _, task := range tasks {
		taskMap[task]++
	}
	taskNodes := make([]*Task, 0)
	for _, cnt := range taskMap {
		taskNodes = append(taskNodes, &Task{remain: cnt, release: 0})
	}
	candidates := Constructor(taskNodes, true)
	blocked := Constructor(make([]*Task, 0), false)
	time := 0
	for candidates.Size() > 0 || blocked.Size() > 0 {
		if candidates.Size() == 0 {
			time = blocked.Top().release
		}
		for blocked.Size() > 0 && blocked.Top().release <= time {
			top := blocked.Pop()
			candidates.Push(top)
		}
		task := candidates.Pop()
		task.release = time + n + 1
		task.remain--
		time++
		if task.remain > 0 {
			blocked.Push(task)
		}
	}
	return time
}