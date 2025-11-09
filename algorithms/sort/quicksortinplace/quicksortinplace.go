// QuickSort In-Place Algorithm (Lomuto partition scheme)
// Time: O(n log n) average, O(n²) worst case
// Space: O(log n) - recursion stack only
// Note: modifies original array

package main

import (
	"fmt"
)

func quickSortInPlace(arr []int, leftIndex, rightIndex int) {
	if leftIndex < rightIndex {
		pivotIndex := partition(arr, leftIndex, rightIndex)
		quickSortInPlace(arr, leftIndex, pivotIndex-1)  // recursively sort left
		quickSortInPlace(arr, pivotIndex+1, rightIndex) // recursively sort right
	}
}

// Lomuto partition scheme
// Places pivot element in correct position and returns its index
func partition(arr []int, leftIndex, rightIndex int) int {
	pivot := arr[rightIndex]
	leftBound := leftIndex - 1

	for i := leftIndex; i < rightIndex; i++ {
		if arr[i] <= pivot {
			leftBound++
			arr[leftBound], arr[i] = arr[i], arr[leftBound]
		}
	}

	arr[leftBound+1], arr[rightIndex] = arr[rightIndex], arr[leftBound+1]
	return leftBound + 1
}

func main() {
	arr := []int{4, 6, 2, 8, 3, 7, 1, 9, 5, 10}
	leftIndex := 0
	rightIndex := len(arr) - 1
	quickSortInPlace(arr, leftIndex, rightIndex)
	fmt.Println(arr)
}

/*
How in-place quick sort works (Lomuto partition):

Example: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]

PARTITION PHASE (in-place rearrangement):
    Initial: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
                                         ↑ pivot = 10 (last element)

    After partition: [4, 6, 2, 8, 3, 7, 1, 9, 5, 10]
                     └──────elements ≤ 10──────┘ │
                     All elements fit, pivot stays at end
                     pivotIndex = 9

    Recursion: quickSort(arr, 0, 8) and quickSort(arr, 10, 9) - nothing

First recursion on [4, 6, 2, 8, 3, 7, 1, 9, 5]:
    pivot = 5 (last)
    After swaps: [4, 2, 3, 1, 5, 7, 6, 9, 8]
                 └──≤5─────┘  │  └──>5────┘
                 pivotIndex = 4

    Recursion: quickSort(arr, 0, 3) and quickSort(arr, 5, 8)

Continue left [4, 2, 3, 1]:
    pivot = 1
    After swaps: [1, 2, 3, 4]
                  │ └─>1───┘
                 pivotIndex = 0

Continue right [7, 6, 9, 8]:
    pivot = 8
    After swaps: [7, 6, 8, 9]
                 └≤8─┘  │  >8
                 pivotIndex = 2

... recursion continues until sorted

RESULT: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] ✓

Key points:
- Pivot selection: LAST element (Lomuto scheme)
- In-place swaps: elements ≤ pivot move left of partition boundary
*/
