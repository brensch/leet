package coin_change

const inf = 1000

// CoinChange returns the minimum number of coins needed to make amount, or -1.
func CoinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	dpArray := []int{}

	for i := 0; i <= amount; i++ {
		minPossible := inf
		for _, coin := range coins {
			if i-coin == 0 {
				minPossible = 1
				break
			}

			if i-coin > 0 && dpArray[i-coin]+1 < minPossible {
				minPossible = dpArray[i-coin] + 1
			}
		}
		dpArray = append(dpArray, minPossible)
	}

	finalAmount := dpArray[amount]
	if finalAmount == inf {
		return -1
	}
	return finalAmount
}
