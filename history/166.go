package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Div          int
	Mod          int
	IndexPastDot int
	Next         *Node
}

func fractionToDecimal(numerator int, denominator int) string {
	// 处理符号
	if numerator == 0 {
		return "0"
	}
	sign := 1
	if numerator < 0 {
		sign = -sign
		numerator = -numerator
	}
	if denominator < 0 {
		sign = -sign
		denominator = -denominator
	}

	var head, tail *Node
	var repeatStart, repeatEnd *Node

	num := numerator
	indexPastDot := 0
	for {
		for num < denominator {
			num *= 10
			indexPastDot++
		}

		div, mod := num/denominator, num%denominator
		node := Node{Div: div, Mod: mod, IndexPastDot: indexPastDot, Next: nil}
		// 初始化链表头
		if head == nil {
			head = &node
		}

		// 添加到链表末尾
		if tail != nil {
			tail.Next = &node
		}
		tail = &node

		if mod == 0 {
			break
		}

		// 检查是否有循环
		repeated := false
		for p := head; p != &node; p = p.Next {
			if p.Mod == mod {
				repeatStart = p
				repeatEnd = &node
				repeated = true
				break
			}
		}
		if repeated {
			break
		}
	}

	result := ""
	if sign < 0 {
		result = "-"
	}

	if head.IndexPastDot == 0 {
		result += strconv.Itoa(head.Div)
		head = head.Next
	} else {
		result += "0"
	}

	currentIndexPastDot := 0
	for p := head; p != nil; p = p.Next {
		if currentIndexPastDot == 0 {
			result += "."
		}

		if p == repeatStart {
			result += "("
		}
		if p == repeatEnd {
			result += ")"
		}

		for currentIndexPastDot < p.IndexPastDot {
			result += "0"
			currentIndexPastDot++
		}
		result += strconv.Itoa(p.Div)
	}
	return result
}

func main() {
	// fmt.Println(fractionToDecimal(1, 6))
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(4, 333))
}
