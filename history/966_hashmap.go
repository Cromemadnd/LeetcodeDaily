package main

import (
	"fmt"
	"strings"
)

func devowel(word string) string {
	var result strings.Builder
	for _, char := range word {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			result.WriteRune('*')
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func spellchecker(wordlist []string, queries []string) []string {
	wordsPerfect := make(map[string]bool)
	wordsCap := make(map[string]string)
	wordsVow := make(map[string]string)

	for _, word := range wordlist {
		wordsPerfect[word] = true
		wordLower := strings.ToLower(word)

		if _, exists := wordsCap[wordLower]; !exists {
			wordsCap[wordLower] = word
		}

		wordDevowel := devowel(wordLower)
		if _, exists := wordsVow[wordDevowel]; !exists {
			wordsVow[wordDevowel] = word
		}
	}

	result := make([]string, len(queries))
	for i, query := range queries {
		if wordsPerfect[query] {
			result[i] = query
			continue
		}
		queryLower := strings.ToLower(query)
		if word, exists := wordsCap[queryLower]; exists {
			result[i] = word
			continue
		}
		queryDevowel := devowel(queryLower)
		if word, exists := wordsVow[queryDevowel]; exists {
			result[i] = word
			continue
		}
		result[i] = ""
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
