// Insertion Sort Algorithm
// Time: O(n²) worst case, O(n) best case
// Space: O(1)

package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		trackValue := arr[i]
		leftIndex := i - 1

		for leftIndex >= 0 && arr[leftIndex] > trackValue {
			arr[leftIndex+1] = arr[leftIndex]
			leftIndex--
		}
		arr[leftIndex+1] = trackValue
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{1, 3, 2, 6, 4, 99, 7}))
}
