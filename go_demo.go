package main

import (
    "fmt"
    "math"
    "time"
)

// Define a struct (similar to a class in other languages)
type Person struct {
    Name string
    Age  int
    City string
}

// Method for the Person struct
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s, %d years old, from %s", p.Name, p.Age, p.City)
}

// Function that demonstrates various Go features
func demonstrateGoFeatures() {
    fmt.Println("=== Go Language Fundamentals Demo ===")

    // Variables and constants
    var message string = "Welcome to Go!"
    const pi = 3.14159
    age := 25 // Short variable declaration

    fmt.Printf("Message: %s\n", message)
    fmt.Printf("Pi: %.2f\n", pi)
    fmt.Printf("Age: %d\n", age)

    // Arrays and slices
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Printf("Numbers: %v\n", numbers)

    // Maps
    colors := map[string]string{
        "red":   "#FF0000",
        "green": "#00FF00",
        "blue":  "#0000FF",
    }

    fmt.Printf("Colors: %v\n", colors)

    // Control structures
    for i, num := range numbers {
        if num%2 == 0 {
            fmt.Printf("Index %d: %d is even\n", i, num)
        } else {
            fmt.Printf("Index %d: %d is odd\n", i, num)
        }
    }

    // Switch statement
    today := time.Now().Weekday()
    switch today {
    case time.Saturday, time.Sunday:
        fmt.Println("It's weekend!")
    default:
        fmt.Println("It's a weekday")
    }
}

// Function demonstrating error handling
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Goroutine function for concurrency demo
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Millisecond * 100)
        results <- job * 2
    }
}

// Main function - entry point of the program
func main() {
    fmt.Println("🚀 Starting Go Programming Demo")
    fmt.Println()

    // Basic features demonstration
    demonstrateGoFeatures()
    fmt.Println()

    // Struct and methods
    person := Person{
        Name: "Alice",
        Age:  30,
        City: "New York",
    }

    fmt.Println("=== Structs and Methods ===")
    fmt.Println(person.Greet())
    fmt.Println()

    // Error handling
    fmt.Println("=== Error Handling ===")
    result, err := divide(10, 2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 / 2 = %.2f\n", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Result: %.2f\n", result)
    }
    fmt.Println()

    // Concurrency with goroutines and channels
    fmt.Println("=== Concurrency with Goroutines ===")

    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for r := 1; r <= 5; r++ {
        result := <-results
        fmt.Printf("Result: %d\n", result)
    }

    fmt.Println()
    fmt.Println("✅ Demo completed successfully!")
}