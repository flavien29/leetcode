package main

import fmt "fmt"

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
  result := make([]int, 2)
  for i := 0; i < len(nums); i++ {
    value, ok := mymap[nums[i]]
    if ok {
      result[0] = value
      result[1] = i
    } else {
      mymap[target - nums[i]] = i
    }
  }
  return result
}

func main() {
	target := 15
	nums := []int{3, 2, 4, 13}
	index := twoSum(nums, target)
	fmt.Printf("%v", index)
}
