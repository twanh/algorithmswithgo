# Quick sort

> Implementing the quick sort algorithm in go

## What is quicksort?
Quicksort is a algorithm that sorts an array (list,slice,map) of (in this case) numbers frow the lowest value to the highest value. 

## How does it work?
Quick sort works with selecting a pivot and then compares the other items in the slice to the pivot.

If the current item is lower then the pivot it sorts it to the left of the pivot, if it is higher it sorts it to the right. 
So you have two sub-arrays, one with the values lower then the pivot and one with values higher then the pivot. We then use the same quicksort (recursively) to sort the sub-arrays.

### Steps:
1. **Create exit condition:** We need this to make sure our `QuickSort` does no infinitely continue, at some point it has finished. 
    In our case it has finished when the `len(input)<=1` because in that case it cannot be sorted anymore. *Because*: the only two options are [], and [x]  
2. **Pivot selection:** Pick an element as the pivot
    We select the last item in the slice (referred to as the last item om the right)
3. **Partitioning:** Reorder the slice (array,map,etc..) so that all the elements that are smaller then the pivot are  left of the pivot and the elements larger then the array are to the right.
   * We do this using a for-loop which loops over all the items in the slice
   * When an item is lower then the pivot (`val < pivot`) we swap it with the item that is currently on the `curLowIndx`. (`curLowIndx` keeps track of what value to replace). After the swap we incease the `curLowIndex` variable with `1`.
   * If an item is higher we just continue the loop  
     --> Notice here that we do not increase the `curLowIndx` so that the item that is higher (than the pivot) will be swapped as soon as we encounter an item that is lower (than the pivot). So at this point the `curLowIndx` points to an item that is actually higher then the pivot.
   * When the loop finishes (has looped over all the items in the slice) we set the `pivot` to the `curLowIndx`. This way the pivot is exactly in the right place.
   
   
4. **Sort sub-arrays:** Recursively apply steps 1, 2 & 3 to the sub-arrays (the values to the left of the pivot and to the right of the pivot)
     * The sub arrays are  `sorted[:curLowIndx]` & `sorted[curLowIndx+1:]`. This way we split the sorted slice (=array) in two parts where the first one are all the numbers smaller then the current pivot and the seccond one are all the numbers larger then the current pivot.  *Note* that the pivot is not part of these sub-arrays (which makes sense ofc) because we do `cuIndex+1:` for the 'right' array.