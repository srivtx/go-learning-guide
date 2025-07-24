package main

import (
	"fmt"
	"sort"
	"strings"
)

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("=== Arrays, Slices, and Maps Comprehensive Guide ===")

	// ==================== ARRAYS ====================
	fmt.Println("\n--- ARRAYS ---")

	// Array declaration and initialization
	var numbers [5]int // Array of 5 integers, zero-initialized
	fmt.Printf("Empty array: %v\n", numbers)

	// Initialize with values
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Printf("Prime numbers: %v\n", primes)

	// Let compiler count elements
	fruits := [...]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits: %v (length: %d)\n", fruits, len(fruits))

	// Sparse initialization (specific indices)
	sparse := [10]int{1: 10, 5: 50, 9: 90}
	fmt.Printf("Sparse array: %v\n", sparse)

	// Multi-dimensional arrays
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Printf("Matrix:\n")
	for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", matrix[i])
	}

	// ==================== SLICES ====================
	fmt.Println("\n--- SLICES ---")

	// Slice literals
	colors := []string{"red", "green", "blue"}
	fmt.Printf("Colors slice: %v (len: %d, cap: %d)\n", colors, len(colors), cap(colors))

	// Creating slices from arrays
	arrayNumbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := arrayNumbers[2:7]   // Elements 2 to 6
	slice2 := arrayNumbers[:5]    // Elements 0 to 4
	slice3 := arrayNumbers[5:]    // Elements 5 to end
	slice4 := arrayNumbers[:]     // All elements

	fmt.Printf("Original array: %v\n", arrayNumbers)
	fmt.Printf("slice[2:7]: %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice[:5]: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice[5:]: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice[:]: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))

	// Make function to create slices
	madeSlice := make([]int, 5)      // length 5, capacity 5
	fmt.Printf("Made slice: %v (len: %d, cap: %d)\n", madeSlice, len(madeSlice), cap(madeSlice))

	madeSliceWithCap := make([]int, 3, 10) // length 3, capacity 10
	fmt.Printf("Made slice with capacity: %v (len: %d, cap: %d)\n", madeSliceWithCap, len(madeSliceWithCap), cap(madeSliceWithCap))

	// Append to slices
	var dynamicSlice []int
	fmt.Printf("Initial dynamic slice: %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	dynamicSlice = append(dynamicSlice, 1)
	fmt.Printf("After append(1): %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	dynamicSlice = append(dynamicSlice, 2, 3, 4)
	fmt.Printf("After append(2,3,4): %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Append another slice
	moreNumbers := []int{5, 6, 7}
	dynamicSlice = append(dynamicSlice, moreNumbers...)
	fmt.Printf("After append slice: %v (len: %d, cap: %d)\n", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	// Copy slices
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	copy(destination, source)
	fmt.Printf("Source: %v, Destination: %v\n", source, destination)

	// Slice tricks
	fmt.Println("\n--- SLICE TRICKS ---")

	// Remove element at index
	numbers2 := []int{10, 20, 30, 40, 50}
	indexToRemove := 2
	numbers2 = append(numbers2[:indexToRemove], numbers2[indexToRemove+1:]...)
	fmt.Printf("After removing index 2: %v\n", numbers2)

	// Insert element at index
	numbers3 := []int{1, 2, 4, 5}
	indexToInsert := 2
	valueToInsert := 3
	numbers3 = append(numbers3[:indexToInsert], append([]int{valueToInsert}, numbers3[indexToInsert:]...)...)
	fmt.Printf("After inserting 3 at index 2: %v\n", numbers3)

	// Filter slice (keep even numbers)
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evens []int
	for _, num := range original {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	fmt.Printf("Even numbers from %v: %v\n", original, evens)

	// ==================== MAPS ====================
	fmt.Println("\n--- MAPS ---")

	// Map literals
	countries := map[string]string{
		"US": "United States",
		"FR": "France",
		"JP": "Japan",
		"BR": "Brazil",
	}
	fmt.Printf("Countries: %v\n", countries)

	// Make function to create maps
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35
	fmt.Printf("Ages: %v\n", ages)

	// Map operations
	fmt.Println("\n--- MAP OPERATIONS ---")

	// Check if key exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Printf("Alice's age: %d\n", age)
	}

	age, exists = ages["David"]
	if !exists {
		fmt.Println("David's age not found")
	}

	// Delete from map
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)

	// Iterate over map
	fmt.Println("Iterating over countries:")
	for code, name := range countries {
		fmt.Printf("  %s: %s\n", code, name)
	}

	// Map with struct values
	people := map[string]Person{
		"p1": {Name: "Alice", Age: 30, City: "New York"},
		"p2": {Name: "Bob", Age: 25, City: "San Francisco"},
		"p3": {Name: "Charlie", Age: 35, City: "Chicago"},
	}

	fmt.Println("\nPeople map:")
	for id, person := range people {
		fmt.Printf("  %s: %+v\n", id, person)
	}

	// ==================== ADVANCED OPERATIONS ====================
	fmt.Println("\n--- ADVANCED OPERATIONS ---")

	// Sorting slices
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Before sorting: %v\n", words)
	sort.Strings(words)
	fmt.Printf("After sorting: %v\n", words)

	nums := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Before sorting: %v\n", nums)
	sort.Ints(nums)
	fmt.Printf("After sorting: %v\n", nums)

	// Custom sorting
	people2 := []Person{
		{Name: "Alice", Age: 30, City: "New York"},
		{Name: "Bob", Age: 25, City: "San Francisco"},
		{Name: "Charlie", Age: 35, City: "Chicago"},
	}

	// Sort by age
	sort.Slice(people2, func(i, j int) bool {
		return people2[i].Age < people2[j].Age
	})
	fmt.Printf("People sorted by age: %v\n", people2)

	// Sort by name
	sort.Slice(people2, func(i, j int) bool {
		return people2[i].Name < people2[j].Name
	})
	fmt.Printf("People sorted by name: %v\n", people2)

	// Working with strings as slices
	text := "Hello, World!"
	fmt.Printf("Original text: %s\n", text)
	fmt.Printf("Text as bytes: %v\n", []byte(text))
	fmt.Printf("Text as runes: %v\n", []rune(text))

	// String manipulation
	sentence := "Go is awesome for programming"
	words2 := strings.Fields(sentence)
	fmt.Printf("Words in sentence: %v\n", words2)

	joined := strings.Join(words2, "-")
	fmt.Printf("Joined with hyphens: %s\n", joined)

	// 2D slice (slice of slices)
	fmt.Println("\n--- 2D SLICES ---")
	
	// Create a 2D slice
	board := make([][]string, 3)
	for i := range board {
		board[i] = make([]string, 3)
		for j := range board[i] {
			board[i][j] = fmt.Sprintf("(%d,%d)", i, j)
		}
	}

	fmt.Println("2D board:")
	for i := range board {
		fmt.Printf("  %v\n", board[i])
	}

	// Map of slices
	fmt.Println("\n--- MAP OF SLICES ---")
	
	grades := map[string][]int{
		"Alice":   {85, 92, 78, 94},
		"Bob":     {76, 89, 91, 87},
		"Charlie": {92, 88, 85, 90},
	}

	for student, scores := range grades {
		total := 0
		for _, score := range scores {
			total += score
		}
		average := float64(total) / float64(len(scores))
		fmt.Printf("%s: scores %v, average %.1f\n", student, scores, average)
	}
}
