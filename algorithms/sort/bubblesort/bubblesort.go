// Bubble Sort Algorithm
// Time Complexity: O(n²) - worst and average case, O(n) - best case (already sorted)
// Space Complexity: O(1) - sorts in-place

package main

import "fmt"

func bubbleSort(arr []int) []int {
	sortedUntilIndex := len(arr) - 1
	sorted := false

	for !sorted {
		sorted = true
		for i := 0; i < sortedUntilIndex; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		sortedUntilIndex--
	}
	return arr
}

func main() {
	fmt.Println(bubbleSort([]int{1, 3, 4, 2, 7, 5, 6, 8}))
}
