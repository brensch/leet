package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {

	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		// {
		// 	name:          "single dependency chain",
		// 	numCourses:    2,
		// 	prerequisites: [][]int{{1, 0}},
		// 	want:          true,
		// },
		// {
		// 	name:          "simple cycle",
		// 	numCourses:    2,
		// 	prerequisites: [][]int{{1, 0}, {0, 1}},
		// 	want:          false,
		// },
		// {
		// 	name:          "multiple disconnected components",
		// 	numCourses:    5,
		// 	prerequisites: [][]int{{1, 0}, {2, 1}, {4, 3}},
		// 	want:          true,
		// },
		// {
		// 	name:          "no prerequisites",
		// 	numCourses:    4,
		// 	prerequisites: nil,
		// 	want:          true,
		// },
		// {
		// 	name:          "shared prerequisite is valid",
		// 	numCourses:    3,
		// 	prerequisites: [][]int{{1, 0}, {2, 0}},
		// 	want:          true,
		// },
		{
			name:          "diamond dependency is valid",
			numCourses:    4,
			prerequisites: [][]int{{3, 1}, {3, 2}, {1, 0}, {2, 0}},
			want:          true,
		},
		// {
		// 	name:          "self cycle",
		// 	numCourses:    1,
		// 	prerequisites: [][]int{{0, 0}},
		// 	want:          false,
		// },
		// {
		// 	name:          "three course cycle",
		// 	numCourses:    3,
		// 	prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
		// 	want:          false,
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CanFinish(tc.numCourses, tc.prerequisites)
			if got != tc.want {
				t.Fatalf("CanFinish(%d, %v) = %t, want %t", tc.numCourses, tc.prerequisites, got, tc.want)
			}
		})
	}
}
