// Factorial Calculation
// Time: O(n) - iterates from n down to 1
// Space: O(1) - only uses constant extra space

package main

import "fmt"

func factorial(n int) int {
	num := 1
	for i := n; i >= 1; i-- {
		num *= i
	}
	return num
}

func main() {
	fmt.Println(factorial(5))
}
