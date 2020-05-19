package sorting

// QuickSort sorts the given slice using the quicksort algorithm
// Returns the sorted slice
func QuickSort(input []int) []int {

	// If the length is lower then 1 or 1, the slice empty or has a length of 1 ([] or [x])
	// in that case it does not require sorting
	if len(input) <= 1 {
		return input
	}

	// We copy the input slice to the sorted slice, because otherwise we would
	// override the data in the input slice (also the variable it is stored in) because of
	// how go handles slices
	// see: https://blog.golang.org/slices-intro
	sorted := make([]int, len(input))
	copy(sorted, input)

	// Select pivot
	// we select the right most item
	pivotIndx := len(sorted) - 1 // The index of the last item of the array
	pivot := sorted[pivotIndx]   // The value of the last item in the array
	// The value that keeps track of where the current value that is smaller then the pivot should be stored
	curLowIndx := 0
	// Sort around the pivot
	for indx, val := range sorted {
		if val < pivot {
			// Swap the items
			sorted[indx], sorted[curLowIndx] = sorted[curLowIndx], sorted[indx]
			// Increment the index on which the low position is stored.
			// We need to do this so that the next item that is lower can be stored/swapped with the correct position
			curLowIndx++
		}
	}

	// Add the pivot to it's correct position
	// We know this is the correct position because everything lower is at :curLowIndx ([0, ..., curLowIndx])
	// and everything should be: [curLowIndx, ...]
	sorted[curLowIndx], sorted[pivotIndx] = sorted[pivotIndx], sorted[curLowIndx]

	// Sort the sub-arrays
	left := QuickSort(sorted[:curLowIndx])
	right := QuickSort(sorted[curLowIndx+1:])

	// Combine the sorted sub-arrays into one array (=slice) we can return
	var ret []int
	ret = append(ret, left...)
	ret = append(ret, sorted[curLowIndx])
	ret = append(ret, right...)

	// Return the sorted slice
	return ret

}
