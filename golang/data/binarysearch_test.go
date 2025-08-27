package data

import (
	"fmt"
	"testing"
)

// worst case o(log n), average case o(log n) and best case is o(1)
func TestBinarySearch(t *testing.T) {
	input := []int{-1, 2, 70, 1000, 2002, 3455}
	key := binarySearch(input, -1)
	fmt.Println(key)
}

// we are passing the value and not the reference, cause the input will be modified,
// make sure the input is a sorted array
func binarySearch(input []int, key int) int {
	// calculate the mid
	low := 0
	high := len(input) - 1
	val := recursiveCall(low, high, input, key)
	return val
}

// this will return the index of the element which is found!
func recursiveCall(low, high int, arr []int, key int) int {
	for low <= high {
		// if low and high are very large numbers there can be a overflow so we are going for low+(high-low)/2
		mid := low + (high-low)/2
		if arr[mid] < key {
			low = mid + 1
			return recursiveCall(low, high, arr, key)

		} else if arr[mid] > key {
			high = mid - 1
			return recursiveCall(low, high, arr, key)
		} else {
			// return the index if found
			return int(mid)
		}
	}
	return -1
}
