package main

import (
    "fmt"
    "strings"
    "unicode"
)

func countWordFrequency(input string) map[string]int {
    // Convert to lowercase and remove punctuation
    cleanedInput := strings.ToLower(input)
    cleanedInput = removePunctuation(cleanedInput)
    
    // Split the cleaned input into words
    words := strings.Fields(cleanedInput)
    
    // Count the frequency of each word
    wordFrequency := make(map[string]int)
    for _, word := range words {
        wordFrequency[word]++
    }

    return wordFrequency
}

// removePunctuation removes punctuation from a string.
func removePunctuation(input string) string {
    var result strings.Builder
    for _, ch := range input {
        if !unicode.IsPunct(ch) && !unicode.IsSpace(ch) {
            result.WriteRune(ch)
        } else if unicode.IsSpace(ch) {
            result.WriteRune(ch)
        }
    }
    return result.String()
}

func isPalindrome(s string) bool {
    // Convert to lowercase and remove non-alphanumeric characters
    cleanedInput := strings.ToLower(s)
    cleanedInput = removeNonAlphanumeric(cleanedInput)
    
    // Check if the cleaned string is a palindrome
    for i := 0; i < len(cleanedInput)/2; i++ {
        if cleanedInput[i] != cleanedInput[len(cleanedInput)-1-i] {
            return false
        }
    }

    return true
}

func removeNonAlphanumeric(input string) string {
    var result strings.Builder
    for _, ch := range input {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            result.WriteRune(ch)
        }
    }
    return result.String()
}

func main() {
    s := "Learning at GOG ta gninrael"
    frequencyMap := countWordFrequency(s)
    fmt.Println(frequencyMap)
    fmt.Println(isPalindrome(s))
}
