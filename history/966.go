package main

import (
	"fmt"
)

func toLowerCase(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func looseMatchLevel(word string, query string) int {
	if word == query {
		return 3 // 完全匹配
	}

	result := 2 // 默认为大小写宽松匹配
	for ind, char := range query {
		if ind >= len(word) { // 表示长度不匹配
			return 0
		}

		loweredChar := toLowerCase(char)
		loweredWordChar := toLowerCase(rune(word[ind]))
		if loweredChar != loweredWordChar { // 表示不符合大小写宽松匹配
			result = 1                                              // 降级为元音宽松匹配
			if !isVowel(loweredChar) || !isVowel(loweredWordChar) { // 必须要二者都是元音，才能算作"宽松匹配"
				return 0
			}
		}
	}
	return result
}

func spellchecker(wordlist []string, queries []string) []string {
	result := make([]string, len(queries))
	for index, query := range queries {
		matchWord := ""
		matchLevel := 0

		for _, word := range wordlist {
			myLevel := looseMatchLevel(word, query)
			if myLevel > matchLevel {
				matchLevel = myLevel
				matchWord = word

				if matchLevel == 3 {
					break
				}
			}
		}

		result[index] = matchWord
	}
	return result
}

func main() {
	// 示例 1
	wordlist1 := []string{"KiTe", "kite", "hare", "Hare"}
	queries1 := []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"}
	result1 := spellchecker(wordlist1, queries1)
	fmt.Printf("示例 1:\n")
	fmt.Printf("wordlist = %v\n", wordlist1)
	fmt.Printf("queries = %v\n", queries1)
	fmt.Printf("输出 = %v\n\n", result1)
	// 期望输出: ["kite","KiTe","KiTe","Hare","hare","","","KiTe","","KiTe"]

	// 示例 2
	wordlist2 := []string{"yellow"}
	queries2 := []string{"YellOw"}
	result2 := spellchecker(wordlist2, queries2)
	fmt.Printf("示例 2:\n")
	fmt.Printf("wordlist = %v\n", wordlist2)
	fmt.Printf("queries = %v\n", queries2)
	fmt.Printf("输出 = %v\n\n", result2)
	// 期望输出: ["yellow"]

	// 自定义测试
	wordlist3 := []string{"apple", "Orange", "BANANA"}
	queries3 := []string{"apple", "orange", "APPLE", "bAnAnA", "ipple", "oronge"}
	result3 := spellchecker(wordlist3, queries3)
	fmt.Printf("自定义测试:\n")
	fmt.Printf("wordlist = %v\n", wordlist3)
	fmt.Printf("queries = %v\n", queries3)
	fmt.Printf("输出 = %v\n\n", result3)
	// 期望输出: ["apple","Orange","apple","BANANA","apple","Orange"]
}
