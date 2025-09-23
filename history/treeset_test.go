package main

import (
	"fmt"

	"github.com/emirpasic/gods/sets/treeset"
)

func main() {
	// 创建一个 TreeSet，使用默认 int 比较器
	set := treeset.NewWithIntComparator()

	// 插入一些数据
	set.Add(1, 3, 5, 7, 9)

	// ----------- 精确查找（相当于二分查找 key 是否存在） -----------
	key := 5
	if set.Contains(key) {
		fmt.Printf("%d 存在于集合中\n", key)
	} else {
		fmt.Printf("%d 不存在于集合中\n", key)
	}

	// ----------- 找到大于等于 key 的最小值 (lower bound) -----------
	key = 4
	if ceil, found := set.Ceiling(key); found {
		fmt.Printf("集合中 >= %d 的最小元素是 %v\n", key, ceil)
	} else {
		fmt.Printf("集合中不存在 >= %d 的元素\n", key)
	}

	// ----------- 找到小于等于 key 的最大值 (upper bound - 1) -----------
	key = 6
	if floor, found := set.Floor(key); found {
		fmt.Printf("集合中 <= %d 的最大元素是 %v\n", key, floor)
	} else {
		fmt.Printf("集合中不存在 <= %d 的元素\n", key)
	}

	// ----------- 遍历 (已排序) -----------
	fmt.Print("TreeSet 中的排序结果: ")
	it := set.Iterator()
	for it.Next() {
		fmt.Print(it.Value(), " ")
	}
	fmt.Println()
}
