# Binary search
> Implementing the binary search algorithm in go 

## What is binary search

Binary search is a searching method that can search a sorted array.
It does this by for every step taking the average of the minimum value of the array and the maximum value of the array. 
If the value that is stored at the index of middle is not equal to the the target (searched) value it checks if the target value is larger or smaller.
Based on this it eliminates a part of the array where the target will not be stored. 

### Steps:
1. Calculate the starting `min` and `max` values. 
     * `min` is `0` because that is lowest index in an array (slice)
     * `max` is `len(input)-1`, which is the last index of the (`input`) array.
2. Loop as long as `min < max`  
   We use this condition for the loop because we increment `min` and decrement `max` so the difference will decrease per step so there is no infinite loop.  
     * Check if the `input[middle]` is equal to the `target` value and `return target` if true
     * Otherwise we can check if the value at the middle index in the input array is smaller then the target value.  
      If so, we can set the minimum value (`min`) to `middle + 1` because we now now that target will be more then the current `middle` value
     * If the value at the middle index is larger then the target we can set the `max` to `middle-1` because we now know that the target is a lower number then the current middle and the current `middle` is the floored avarage of the `min` and `max` so most definitely smaller then the `max`. Se we decrease the `max` with the `middle-1`.
  3. If we exit the loop without returning in the loop this meaans that the target value was not found and we return `-1`.