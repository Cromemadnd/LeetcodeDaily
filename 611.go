package main

import (
	"fmt"
	"sort"
)

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	n := len(nums)
	sort.Ints(nums)

	ans := 0
	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			if v+nums[j] > nums[j] {
				ans += sort.SearchInts(nums[j+1:], v+nums[j])
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{0, 1, 1, 1}))
	fmt.Println(triangleNumber([]int{1, 2, 3, 4}))
}
