package main

import "fmt"

func triangleNumber(nums []int) int {
	bucket := [1001]int{}
	for _, num := range nums {
		bucket[num]++
	}

	numsUnique := []int{}
	for i := 1; i <= 1000; i++ {
		if bucket[i] != 0 {
			numsUnique = append(numsUnique, i)
		}
	}

	fmt.Println(numsUnique)

	count := 0
	for iIndex, i := range numsUnique {
		for jIndex, j := range numsUnique[iIndex:] {
			for _, k := range numsUnique[iIndex+jIndex:] {
				if i+j <= k {
					break
				}

				if i == j {
					if j == k {
						count += bucket[i] * (bucket[i] - 1) * (bucket[i] - 2) / 6
					} else {
						count += bucket[i] * (bucket[i] - 1) * bucket[k] / 2
					}
				} else {
					if j == k {
						count += bucket[i] * bucket[j] * (bucket[j] - 1) / 2
					} else if i == k {
						count += bucket[i] * bucket[j] * (bucket[i] - 1) / 2
					} else {
						count += bucket[i] * bucket[j] * bucket[k]
					}
				}
				// fmt.Println(i, j, k, iIndex, jIndex, kIndex, count)
			}
		}
	}
	return count
}

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{4, 2, 3, 4}))
	fmt.Println(triangleNumber([]int{0, 1, 1, 1}))
	fmt.Println(triangleNumber([]int{1, 2, 3, 4}))
}
