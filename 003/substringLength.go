package main

import fmt "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	mymap := map[byte]int{}
	for i := range s {
		oldIndex, ok := mymap[s[i]]
		if !ok {
			mymap[s[i]] = i
			if len(mymap) > max {
				max = len(mymap)
			}
		} else {
			mymap[s[i]] = i
			for key, val := range mymap {
				if val < oldIndex {
					delete(mymap, key)
				}
			}
		}
	}
	return max
}

func main() {
	var str string = "totorto"
	length := lengthOfLongestSubstring(str)
	fmt.Printf("%v", length)
}
