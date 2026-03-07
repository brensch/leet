package binary_search

import "fmt"

// Search looks for target in a sorted slice.
func Search(nums []int, target int) int {

	// int rounds, potential source of error
	index := len(nums) / 2
	min := 0
	max := len(nums) - 1

	badTimes := 10

	for i := range badTimes {
		fmt.Println(i, min, max, target, index)
		current := nums[index]
		if current == target {
			return index
		}
		// search down
		if current > target {
			max = index
		} else {
			min = index
		}

		index = min + (max-min)/2
		if (max-min)/2 == 0 {
			break
		}
	}
	return -1
}
