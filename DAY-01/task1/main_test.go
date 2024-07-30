package main

import (
    "testing"
)

func TestCalculateAverage(t *testing.T) {
    grades := map[string]int{
        "Math":      90,
        "English":   85,
        "Science":   80,
        "History":   75,
        "Geography": 92,
    }
    expected := 84.4
    actual := calculateAverage(grades)
    if actual != expected {
        t.Errorf("calculateAverage() = %.2f; expected %.2f", actual, expected)
    }
}
