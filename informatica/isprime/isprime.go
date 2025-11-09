// Prime Number Check
// Time: O(√n) - checks divisors up to square root of n
// Space: O(1) - only uses constant extra space

package main

import (
	"fmt"
)

func IsPrimeNumber(n int) (int, bool) {
	if n <= 1 {
		return n, false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return n, false
		}
	}
	return n, true
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(IsPrimeNumber(n))
}
