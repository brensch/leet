package top_k_frequent_elements

import "fmt"

// TopKFrequent returns the k most frequent values in nums.
func TopKFrequent(nums []int, k int) []int {

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	sortedTallies := make([]Tally, k)

	for val, count := range counts {
		for i, tally := range sortedTallies {
			if count > tally.Count {
				// push all down, skipping last (gets pushed out)
				for j := i; j < len(sortedTallies)-1; j++ {
					sortedTallies[j+1] = sortedTallies[j]
				}
				sortedTallies[i] = Tally{Val: val, Count: count}
				break
			}
		}
	}
	fmt.Println(sortedTallies)

	finalNums := make([]int, 0, k)
	for _, tally := range sortedTallies {
		finalNums = append(finalNums, tally.Val)
	}

	return finalNums
}

type Tally struct {
	Val   int
	Count int
}
