package main

import (
	"fmt"
)

func reverse(x int) int {
	maxint32 := 1<<31 - 1
	newVal := 0
	negativeValue := false

	if x < 0 {
		negativeValue = true
		x = -x
	}
	for x > 0 {
		number := x % 10
		newVal *= 10
		newVal += number
		x /= 10
	}
	if newVal > maxint32 {
		return 0
	}
	if negativeValue {
		return -newVal
	}

	return newVal
}

func main() {
	var val int = 12345678
	res := reverse(val)
	fmt.Printf("%v", res)
}
