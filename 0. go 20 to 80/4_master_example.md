Below is a master example code that demonstrates all the common keywords, rules, and symbols in Go, with explanations in Markdown format.

---

# Go Master Example Code

```go
// main.go
package main // Declares the package; "main" is the entry point.

import (
    "fmt"        // Importing standard libraries
    "math"       // Importing math for constants like Pi
    "errors"     // Importing errors to handle errors
)

// Constants
const Pi = 3.14

// Global variable declaration
var globalVar = "I am accessible everywhere in the main package"

// Custom struct declaration
type Person struct {
    Name string
    Age  int
}

// Interface declaration
type Greeter interface {
    Greet() string
}

// Constructor-like function to initialize a new Person (not a true constructor, as Go doesn't have constructors)
func NewPerson(name string, age int) *Person {
    return &Person{Name: name, Age: age}
}

// Method for Person struct that implements the Greeter interface
func (p *Person) Greet() string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

// Basic function with a named return and error handling
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

// Variadic function that takes a variable number of arguments
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

// Function demonstrating pointer usage
func increment(num *int) {
    *num = *num + 1
}

// init function for setup
func init() {
    fmt.Println("Initializing the program...")
}

// Main function - entry point of the application
func main() {
    fmt.Println("Welcome to the Go Master Example Code!")

    // Variable declarations
    var localVar int = 42      // Declaring a variable with an explicit type
    anotherVar := "Quick init" // Short declaration with type inference

    fmt.Println("Global Variable:", globalVar)
    fmt.Println("Local Variable:", localVar)
    fmt.Println("Another Variable:", anotherVar)

    // Working with constants
    fmt.Printf("The value of Pi is approximately %.2f\n", Pi)
    fmt.Printf("The value of Pi from math package is %.2f\n", math.Pi)

    // Structs and Interfaces
    person := NewPerson("Alice", 30)
    fmt.Println(person.Greet())

    // Function with error handling
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 divided by 2 is %.2f\n", result)
    }

    // Goroutine and channel demonstration
    c := make(chan int)
    go func() {
        c <- sum(1, 2, 3, 4, 5)
    }()
    fmt.Println("Sum from Goroutine:", <-c)

    // Pointers
    value := 10
    increment(&value)
    fmt.Println("Incremented Value:", value)

    // Defer demonstration
    defer fmt.Println("This will run at the end of main")

    // Type assertion with empty interface
    var anyType interface{} = "I am a string"
    str, ok := anyType.(string)
    if ok {
        fmt.Println("Value:", str)
    }

    // Looping structures
    for i := 0; i < 3; i++ {
        fmt.Println("Loop iteration:", i)
    }

    // Switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of the week!")
    case "Friday":
        fmt.Println("Almost weekend!")
    default:
        fmt.Println("A regular day.")
    }
}
```

---

# Explanation of Go Keywords and Rules

### 1. **Package**:

- `package main`: Defines the package. The `main` package indicates the entry point of the application.

### 2. **Imports**:

- Importing standard libraries like `fmt`, `math`, and `errors`.

### 3. **Constants and Variables**:

- `const Pi = 3.14`: A constant declaration. Constants cannot be modified after declaration.
- `var globalVar = "..."`: Global variable accessible throughout the `main` package.
- Local variable declaration: `var localVar int = 42` (explicit type) and `anotherVar := "..."` (short variable declaration with type inference).

### 4. **Structs and Interfaces**:

- `type Person struct {...}`: Defines a struct to group related data.
- `type Greeter interface {...}`: Defines an interface that `Person` implements by defining `Greet()`.

### 5. **Functions**:

- `func divide(a, b float64) (result float64, err error)`: A function with named return values and error handling.
- `func sum(nums ...int)`: A variadic function that can accept multiple arguments.
- `func increment(num *int)`: Demonstrates pointers, allowing changes to the original variable.

### 6. **Concurrency**:

- `go func() {...}()`: Starts a goroutine, running code concurrently.
- `c := make(chan int)`: Creates a channel to communicate between goroutines.

### 7. **Control Structures**:

- **For Loop**: `for i := 0; i < 3; i++ {...}` - The only looping structure in Go.
- **Switch Statement**: `switch day {...}` - Conditional branching without `break` for each case.
- **If Statement**: Demonstrated in error handling.

### 8. **Defer**:

- `defer fmt.Println("...")`: Schedules a function to run after the surrounding function completes.

### 9. **Type Assertions**:

- `var anyType interface{} = "..."`: The empty interface can hold any type.
- `str, ok := anyType.(string)`: Type assertion to convert `interface{}` to a specific type if possible.

### 10. **Special Functions**:

- `func init() {...}`: A special function that runs automatically before `main()` for setup tasks.

This example demonstrates all essential Go keywords, rules, and symbols in a structured and meaningful way, covering everything from package basics to concurrency and error handling.
