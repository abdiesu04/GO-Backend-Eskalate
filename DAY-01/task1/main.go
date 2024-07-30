// main.go
package main

import (
    "fmt"
)

func calculateAverage(grades map[string]int) float64 {
    sum := 0
    for _, value := range grades {
        sum += value
    }
    return float64(sum) / float64(len(grades))
}

func main() {
    fmt.Println("Enter your name: ")
    var name string
    fmt.Scanln(&name)
    fmt.Println("How many Subjects do you take? ")
    var subjects int
    fmt.Scanln(&subjects)
    grades := make(map[string]int)
    for i := 0; i < subjects; i++ {
        fmt.Println("Enter the subject name: ")
        var subject string
        fmt.Scanln(&subject)
        fmt.Println("Enter the marks of the subject:", subject)
        var marks int
        fmt.Scanln(&marks)
        if marks > 100 || marks < 0 {
            fmt.Println("Marks should be between 0 and 100")
            i--
            continue
        }
        grades[subject] = marks
    }
    average := calculateAverage(grades)
    fmt.Printf("Student Name: %s\nGrades: %v\nAverage: %.2f\n", name, grades, average)
}
