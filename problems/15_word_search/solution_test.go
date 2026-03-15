package word_search

import "testing"

func TestExist(t *testing.T) {
	t.Skip("remove this skip when you are ready to solve Word Search")

	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "word exists",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "word requires reusing a cell",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
		{
			name:  "single letter",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Exist(tc.board, tc.word)
			if got != tc.want {
				t.Fatalf("Exist(..., %q) = %t, want %t", tc.word, got, tc.want)
			}
		})
	}
}
