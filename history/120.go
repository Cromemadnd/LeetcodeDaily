package main

import (
	"fmt"
	"math"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := [][]int{{triangle[0][0]}}

	for i := 1; i < n; i++ {
		dp = append(dp, []int{})
		for j := 0; j <= i; j++ {
			if j == 0 {
				dp[i] = append(dp[i], dp[i-1][0]+triangle[i][j])
				continue
			}
			if j == i {
				dp[i] = append(dp[i], dp[i-1][j-1]+triangle[i][j])
				continue
			}
			dp[i] = append(dp[i], min(dp[i-1][j-1], dp[i-1][j])+triangle[i][j])
		}
	}

	result := math.MaxInt
	for _, value := range dp[n-1] {
		if value < result {
			result = value
		}
	}
	return result
}

func main() {
	fmt.Printf("%v", minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
