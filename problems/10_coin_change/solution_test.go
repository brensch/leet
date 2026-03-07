package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{name: "standard case", coins: []int{1, 2, 5}, amount: 11, want: 3},
		{name: "impossible", coins: []int{2}, amount: 3, want: -1},
		{name: "zero amount", coins: []int{1}, amount: 0, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CoinChange(tc.coins, tc.amount)
			if got != tc.want {
				t.Fatalf("CoinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
			}
		})
	}
}
