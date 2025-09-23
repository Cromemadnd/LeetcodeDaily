package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func replaceNonCoprimes(nums []int) []int {
	st := nums[:0] // 把 nums 当作栈用
	for _, x := range nums {
		for len(st) > 0 && gcd(x, st[len(st)-1]) > 1 {
			x = lcm(x, st[len(st)-1])
			st = st[:len(st)-1]
		}
		st = append(st, x)
	}
	return st
}

func main() {
	// 示例 1
	nums1 := []int{6, 4, 3, 2, 7, 6, 2}
	result1 := replaceNonCoprimes(nums1)
	fmt.Printf("输入: %v\n", nums1)
	fmt.Printf("输出: %v\n", result1) // 预期输出: [12, 7, 6]
	fmt.Println("---")

	// 示例 2
	nums2 := []int{2, 2, 1, 1, 3, 3, 3}
	result2 := replaceNonCoprimes(nums2)
	fmt.Printf("输入: %v\n", nums2)
	fmt.Printf("输出: %v\n", result2) // 预期输出: [2, 1, 1, 3]
	fmt.Println("---")

	// 更多测试用例
	nums3 := []int{7, 5, 2, 3, 8}
	result3 := replaceNonCoprimes(nums3)
	fmt.Printf("输入: %v\n", nums3)
	fmt.Printf("输出: %v\n", result3) // 预期输出: [7, 5, 2, 3, 8] (没有非互质数)
	fmt.Println("---")

	nums4 := []int{2, 6, 4, 8}
	result4 := replaceNonCoprimes(nums4)
	fmt.Printf("输入: %v\n", nums4)
	fmt.Printf("输出: %v\n", result4) // 预期输出: [24]
}
