package main

import "fmt"

func maxFreqSum(s string) int {
	dict := make(map[rune]int)
	for _, char := range s {
		dict[char]++
	}

	maxV, maxNV := 0, 0
	for char, val := range dict {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			if val > maxV {
				maxV = val
			}
		default:
			if val > maxNV {
				maxNV = val
			}
		}
	}

	return maxV + maxNV
}

func main() {
	// 示例 1
	s1 := "successes"
	fmt.Printf("输入: %s, 输出: %d\n", s1, maxFreqSum(s1))

	// 示例 2
	s2 := "aeiaeia"
	fmt.Printf("输入: %s, 输出: %d\n", s2, maxFreqSum(s2))

	// 示例 3
	s3 := "bcdf"
	fmt.Printf("输入: %s, 输出: %d\n", s3, maxFreqSum(s3))
}
