type Course struct {
	children []int
	preNum   int
}

func NewCourse() *Course {
	return &Course{
		children: make([]int, 0),
		preNum:   0,
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([]*Course, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		graph = append(graph, NewCourse())
	}
	for _, prerequire := range prerequisites {
		t, s := prerequire[0], prerequire[1]
		graph[s].children = append(graph[s].children, t)
		graph[t].preNum++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if graph[i].preNum == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		for _, child := range graph[course].children {
			graph[child].preNum--
			if graph[child].preNum == 0 {
				queue = append(queue, child)
			}
		}
	}
	for i := 0; i < numCourses; i++ {
		if graph[i].preNum > 0 {
			return false
		}
	}
	return true
}