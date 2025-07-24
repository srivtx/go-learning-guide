package main

import (
	"fmt"
	"math/cmplx"
	"reflect"
)

// Global variables
var globalString = "I'm global!"
var globalInt int // Zero value: 0

// Multiple variable declaration
var (
	name     = "Go Programming"
	version  = 1.21
	isActive = true
)

// Constants
const Pi = 3.14159
const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusPending  = "pending"
)

// Iota for enumeration
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
	Thursday      // 4
	Friday        // 5
	Saturday      // 6
)

func main() {
	fmt.Println("=== Go Basic Types Demo ===")

	// Integer types
	var age int = 25
	var population int64 = 7_800_000_000 // Underscores for readability
	var byteValue byte = 255             // byte is alias for uint8

	fmt.Printf("Age: %d (type: %T)\n", age, age)
	fmt.Printf("Population: %d (type: %T)\n", population, population)
	fmt.Printf("Byte value: %d (type: %T)\n", byteValue, byteValue)

	// Floating point types
	var price float32 = 99.99
	var distance float64 = 12345.6789

	fmt.Printf("Price: %.2f (type: %T)\n", price, price)
	fmt.Printf("Distance: %.4f (type: %T)\n", distance, distance)

	// Complex numbers
	var complex1 complex64 = 1 + 2i
	var complex2 complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Printf("Complex1: %.2f (type: %T)\n", complex1, complex1)
	fmt.Printf("Complex2: %.2f (type: %T)\n", complex2, complex2)

	// Boolean
	var isStudent bool = true
	var isEmployed bool // Zero value: false

	fmt.Printf("Is student: %v, Is employed: %v\n", isStudent, isEmployed)

	// String
	var greeting string = "Hello, World!"
	var empty string // Zero value: ""

	fmt.Printf("Greeting: %q, Empty: %q\n", greeting, empty)

	// Runes (Unicode code points)
	var runeValue rune = '🐹' // rune is alias for int32
	fmt.Printf("Rune: %c (Unicode: %U, type: %T)\n", runeValue, runeValue, runeValue)

	// Type inference with :=
	city := "New York"        // string
	temperature := 23.5       // float64
	humid := true             // bool
	unicodeChar := 'A'        // rune
	year := 2024              // int

	fmt.Println("\n=== Type Inference ===")
	fmt.Printf("City: %s (%T)\n", city, city)
	fmt.Printf("Temperature: %.1f (%T)\n", temperature, temperature)
	fmt.Printf("Humid: %v (%T)\n", humid, humid)
	fmt.Printf("Unicode char: %c (%T)\n", unicodeChar, unicodeChar)
	fmt.Printf("Year: %d (%T)\n", year, year)

	// Type conversion
	fmt.Println("\n=== Type Conversion ===")
	var x int = 42
	var y float64 = float64(x) // Explicit conversion required
	var z int = int(y)

	fmt.Printf("int to float64: %d -> %.1f\n", x, y)
	fmt.Printf("float64 to int: %.1f -> %d\n", y, z)

	// String to byte slice and back
	message := "Hello, Go!"
	bytes := []byte(message)
	backToString := string(bytes)

	fmt.Printf("String to bytes: %q -> %v\n", message, bytes)
	fmt.Printf("Bytes to string: %v -> %q\n", bytes, backToString)

	// Zero values demonstration
	fmt.Println("\n=== Zero Values ===")
	var (
		zeroInt     int
		zeroFloat   float64
		zeroBool    bool
		zeroString  string
		zeroPointer *int
		zeroSlice   []int
		zeroMap     map[string]int
		zeroFunc    func()
	)

	fmt.Printf("int: %d\n", zeroInt)
	fmt.Printf("float64: %g\n", zeroFloat)
	fmt.Printf("bool: %v\n", zeroBool)
	fmt.Printf("string: %q\n", zeroString)
	fmt.Printf("pointer: %v\n", zeroPointer)
	fmt.Printf("slice: %v (len: %d, cap: %d)\n", zeroSlice, len(zeroSlice), cap(zeroSlice))
	fmt.Printf("map: %v\n", zeroMap)
	fmt.Printf("function: %v\n", zeroFunc)

	// Using reflect to get type information
	fmt.Println("\n=== Type Information ===")
	values := []interface{}{42, 3.14, "hello", true, []int{1, 2, 3}}

	for _, v := range values {
		t := reflect.TypeOf(v)
		fmt.Printf("Value: %v, Type: %v, Kind: %v\n", v, t, t.Kind())
	}

	// Constants demonstration
	fmt.Println("\n=== Constants ===")
	fmt.Printf("Pi: %g\n", Pi)
	fmt.Printf("Status: %s\n", StatusActive)
	fmt.Printf("Today is weekday %d\n", Wednesday)

	// Type-safe constants
	const typedInt int = 100
	const untypedInt = 200

	var a int = typedInt
	var b int64 = untypedInt // Untyped constant can be assigned to compatible types
	var c float64 = untypedInt

	fmt.Printf("Typed constant: %d\n", a)
	fmt.Printf("Untyped constant as int64: %d\n", b)
	fmt.Printf("Untyped constant as float64: %g\n", c)
}
