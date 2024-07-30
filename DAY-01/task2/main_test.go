package main

import (
    "testing"
)

func TestCountWordFrequency(t *testing.T) {
    input := "Go is a statically typed language. Go is Go."
    expected := map[string]int{
        "go":         3,
        "is":         2,
        "a":          1,
        "statically": 1,
        "typed":      1,
        "language":   1,
    }
    actual := countWordFrequency(input)

    for word, count := range expected {
        if actual[word] != count {
            t.Errorf("countWordFrequency(%q) = %d; expected %d", word, actual[word], count)
        }
    }

    for word, count := range actual {
        if expected[word] != count {
            t.Errorf("countWordFrequency() unexpected count for %q: got %d; want %d", word, count, expected[word])
        }
    }
}

func TestIsPalindrome(t *testing.T) {
    cases := []struct {
        input    string
        expected bool
    }{
        {"A man a plan a canal Panama", true},
        {"Go hang a salami I'm a lasagna hog", true},
        {"Not a palindrome", false},
        {"", true},
        {"Hello, World!", false},
        {"Madam, in Eden, I'm Adam", true},
    }

    for _, c := range cases {
        actual := isPalindrome(c.input)
        if actual != c.expected {
            t.Errorf("isPalindrome(%q) = %v; expected %v", c.input, actual, c.expected)
        }
    }
}
