package search

import (
	"math"
)

// BinarySearch searches a (sorted) array and returns the index at which the target is found
func BinarySearch(input []int, target int) int {
	min, max := 0, len(input)-1
	for min < max {
		middle := int(math.Floor(float64((min + max) / 2)))
		if input[middle] == target {
			return middle
		} else if input[middle] < target {
			min = middle + 1
		} else {
			// input[middle] > target
			max = middle - 1
		}
	}
	return -1
}
