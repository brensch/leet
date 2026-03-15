package three_sum

import "slices"

// ThreeSum returns all unique triplets that sum to zero.
func ThreeSum(nums []int) [][]int {

	var trips [][]int

	slices.Sort(nums)

	for i := range len(nums) - 1 {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {

			result := nums[start] + nums[end]

			if result == target {
				trips = append(trips, []int{nums[i], nums[start], nums[end]})
				end--
				start++
				for end > start && nums[end] == nums[end+1] {
					end--
				}
				for end > start && nums[start] == nums[start-1] {
					start++
				}
			}

			if result > target {
				end--
			}

			if result < target {
				start++
			}
		}
	}

	return trips
}
