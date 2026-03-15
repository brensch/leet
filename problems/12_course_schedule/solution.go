package course_schedule

// CanFinish returns whether all courses can be completed.
func CanFinish(numCourses int, prerequisites [][]int) bool {

	// collect paths for each course
	graph := make(map[int][]int)
	for _, prereq := range prerequisites {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	possibleCourses := make(map[int]struct{})
	for i := 0; i < numCourses; i++ {
		visited := make(map[int]struct{})
		_, ok := visited[i]
		if ok {
			continue
		}
		coursePossible := checkCourse(graph, i, visited, possibleCourses)
		if !coursePossible {
			return false
		}
		possibleCourses[i] = struct{}{}
	}
	return true
}

func checkCourse(graph map[int][]int, course int, visited map[int]struct{}, possible map[int]struct{}) bool {
	next, ok := graph[course]
	if !ok {
		return true
	}

	_, isPossible := possible[course]
	if isPossible {
		return true
	}

	for _, nextCourse := range next {

		// if we get to a node and it's already been visited, cooked
		_, ok := visited[nextCourse]
		if ok {
			return false
		}

		visited[nextCourse] = struct{}{}

		if !checkCourse(graph, nextCourse, visited, possible) {
			return false
		}

	}

	return true
}
