// Time: O(n) - single pass through string
// Space: O(n) - new string created

package main

import (
	"fmt"
	"strings"
)

func removeWhitespaces(sentence string) string {
	return strings.ReplaceAll(sentence, " ", "")
}

func main() {
	fmt.Println(removeWhitespaces("From this sentence"))
}

// Time: O(n) - single pass through string
// Space: O(n) - Builder pre-allocates buffer, final string

// func removeWhitespaces(sentence string) string {
// 	var clearStr strings.Builder
// 	clearStr.Grow(len(sentence)) // pre-allocate capacity

// 	for _, char := range sentence {
// 		if char != ' ' {
// 			clearStr.WriteRune(char)
// 		}
// 	}
// 	return clearStr.String()
// }

// Time: O(n²) - each += creates new string and copies all previous chars
// Space: O(n) - result string size proportional to input

// func removeWhitespaces(sentence string) string {
// 	var clearStr string

// 	for i := 0; i < len(sentence); i++ {
// 		if string(sentence[i]) != " " {
// 			clearStr += string(sentence[i])
// 		}
// 	}
// 	return clearStr
// }
