package main

import (
	"fmt"
)

func isPalindromeUtil(x int, y *int) bool {
	if x >= 0 && x < 10 {
		return (x == *y%10)
	}
	if !isPalindromeUtil(x/10, y) {
		return false
	}

	*y /= 10

	return (x%10 == *y%10)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x2 := x

	return isPalindromeUtil(x, &x2)
}

func main() {
	number := 2147447412
	res := isPalindrome(number)
	fmt.Printf("%v", res)
}
