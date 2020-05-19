package main

import (
	"fmt"

	"github.com/twanh/algorithmswithgo/search"
	"github.com/twanh/algorithmswithgo/sorting"
)

func main() {

	fmt.Println("###### SORTING ######")
	toSort := []int{100, 20, 70, 30, 90, 40, 120, 123, 10, 23}
	shouldBeSorted := sorting.QuickSort(toSort)
	fmt.Println("---UNSORTED---")
	fmt.Println(toSort)
	fmt.Println("---SORTED USING QUICKSORT---")
	fmt.Println(shouldBeSorted)
	fmt.Println("###### SEARCHING ######")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 18, 19, 20}
	indxOf12 := search.BinarySearch(numbers, 12)
	fmt.Println("Searching 12 in:", numbers)
	fmt.Println("12 is at the:", indxOf12, "position")

}
