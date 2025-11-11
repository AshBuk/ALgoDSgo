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

/*
How insertion sort works:

Example: [1, 3, 2, 6, 4, 99, 7]

ITERATION 1 (i=1, trackValue=3):
    [1, 3, 2, 6, 4, 99, 7]
     ↑ sorted (leftIndex=0)
        ↑ current (i=1)
    Compare: 1 ≤ 3 ✓ → insert 3 at position 1
    [1, 3, 2, 6, 4, 99, 7]  ← already in correct position

ITERATION 2 (i=2, trackValue=2):
    [1, 3, 2, 6, 4, 99, 7]
        ↑ sorted (leftIndex=1)
           ↑ current (i=2)
    Compare: 3 > 2 ✗ → shift 3 right
    [1, _, 3, 6, 4, 99, 7]
    Compare: 1 ≤ 2 ✓ → insert 2 at position 1
    [1, 2, 3, 6, 4, 99, 7]

ITERATION 3 (i=3, trackValue=6):
    [1, 2, 3, 6, 4, 99, 7]
           ↑ sorted (leftIndex=2)
              ↑ current (i=3)
    Compare: 3 ≤ 6 ✓ → insert 6 at position 3
    [1, 2, 3, 6, 4, 99, 7]  ← already in correct position

	and so on...

Key points:
- Builds sorted portion from left to right
- For each element, shifts larger elements right to make space
- Inserts current element into correct position in sorted portion
- Like sorting a hand of cards: pick card, insert in correct position
- Best case O(n) when array is already sorted (no shifts needed)
- Efficient for small arrays or nearly sorted data
*/
