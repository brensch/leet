package number_of_islands

import "testing"

func cloneGrid(grid [][]byte) [][]byte {
	out := make([][]byte, len(grid))
	for i := range grid {
		out[i] = append([]byte(nil), grid[i]...)
	}
	return out
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "single island",
			grid: [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}},
			want: 1,
		},
		{
			name: "multiple islands",
			grid: [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}},
			want: 3,
		},
		{name: "empty grid", grid: nil, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NumIslands(cloneGrid(tc.grid))
			if got != tc.want {
				t.Fatalf("NumIslands(...) = %d, want %d", got, tc.want)
			}
		})
	}
}
