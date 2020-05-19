package main

import (
	"fmt"

	"github.com/twanh/algorithmswithgo/sorting"
)

func main() {

	// Sorting examples
	toSort := []int{100, 20, 70, 30, 90, 40, 120, 123, 10, 23}

	shouldBeSorted := sorting.QuickSort(toSort)
	fmt.Println("---UNSORTED---")
	fmt.Println(toSort)
	fmt.Println("---SORTED USING QUICKSORT---")
	fmt.Println(shouldBeSorted)

}
