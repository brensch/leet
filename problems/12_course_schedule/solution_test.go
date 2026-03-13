package course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	t.Skip("remove this skip when you are ready to solve Course Schedule")

	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "single dependency chain",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "simple cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "multiple disconnected components",
			numCourses:    5,
			prerequisites: [][]int{{1, 0}, {2, 1}, {4, 3}},
			want:          true,
		},
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
