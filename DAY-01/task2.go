package main

import (
    "fmt"
    "strings"
)

func countWordFrequency(input string) map[string]int {
    cleanedInput := strings.ToLower(strings.ReplaceAll(input, ",", ""))
    words := strings.Split(cleanedInput, " ")

    wordFrequency := make(map[string]int)
    for _, word := range words {
        wordFrequency[word]++
    }

    return wordFrequency
}

func isPalindrome(s string) bool {
	cleanedInput := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), ",", ""))

	for i := 0; i < len(cleanedInput)/2; i++ {
		if cleanedInput[i] != cleanedInput[len(cleanedInput)-1-i] {
			return false
		}
	}

	return true
}

func main() {
    s := "I am learning Go programming language. Go is a statically typed language."
    frequencyMap := countWordFrequency(s)
    fmt.Println(frequencyMap)
    fmt.Println(isPalindrome((s)))
}