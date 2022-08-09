// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// bytes.Buffer version
func comma(s string) string {
	last := strings.LastIndex(s, ".")
	n := last
	if last < 0 {
		n = len(s)
	}

	var newString = bytes.Buffer{}
	for i := 0; i < n; i++ {
		stringLength := len(s[i:n])
		if (stringLength % 3 == 0 && stringLength < n) {
			newString.WriteByte(',')
		}
		newString.WriteByte(s[i])
	}
	if (last < 0) {
		return newString.String()
	}
	return newString.String() + s[last:]
}

//!+
// comma inserts commas in a non-negative decimal integer string.
// recursive version
//func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	return comma(s[:n-3]) + "," + s[n-3:]
//}

//!+
// comma inserts commas in a non-negative decimal integer string.
// string version
//func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	newString := s[n-1:]
//	fmt.Println(newString)
//	for i := len(s) - 2; i >= 0; i-- {
//		newString = string(s[i]) + newString
//		if (len(s[i:]) % 3 == 0 && i > 0) {
//			newString = "," + newString
//		}
//	}
//	return newString
//}

//!-
