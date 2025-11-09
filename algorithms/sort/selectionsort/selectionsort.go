// Selection Sort Algorithm
// Time Complexity: O(n²) - nested loops iterate through the array
// Space Complexity: O(1) - only uses constant extra space for variables

package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minElement := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minElement] {
				minElement = j
			}
		}
		if minElement != i { // swap only if min. element isn't already in correct position
			arr[i], arr[minElement] = arr[minElement], arr[i]
		}
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{1, 6, 5, 4, 7, 2, 3, 8}))
}
