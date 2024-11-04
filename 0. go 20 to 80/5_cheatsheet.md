# Go Language Cheat Sheet

## 1. Go Basics

### 1.1 Package Declaration

- Every Go file starts with a `package` declaration.
- `package main` indicates an entry point, requiring a `main` function to execute.

```go
package main
```

### 1.2 Importing Packages

- Use `import` to include libraries.

```go
import (
    "fmt"
    "math"
    "errors"
)
```

## 2. Variables and Constants

### 2.1 Declaring Variables

- **Explicit declaration**: `var x int = 42`
- **Type inference**: `var x = 42`
- **Short variable declaration** (inside functions only): `x := 42`

```go
var name string = "Go"
age := 21 // Type inferred as int
```

### 2.2 Declaring Constants

- Use `const` for values that do not change.

```go
const Pi = 3.14
const (
    SpeedOfLight = 299792458 // m/s
    EarthGravity = 9.8       // m/s²
)
```

## 3. Data Types

### 3.1 Primitive Types

- Integers: `int`, `int8`, `int16`, `int32`, `int64`
- Floating points: `float32`, `float64`
- Other types: `bool`, `string`, `byte`, `rune`

### 3.2 Composite Types

- **Array**: Fixed-length, homogeneous
  ```go
  var arr [3]int = [3]int{1, 2, 3}
  ```
- **Slice**: Dynamic-length, homogeneous
  ```go
  slice := []int{1, 2, 3, 4}
  ```
- **Map**: Key-value pairs
  ```go
  m := map[string]int{"Alice": 23, "Bob": 25}
  ```

## 4. Control Structures

### 4.1 If-Else

- Conditions do not require parentheses.
- Initialize variables inline with conditions.

```go
if age := 21; age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### 4.2 Switch

- Go’s `switch` automatically breaks after each case.
- `default` handles all unmatched cases.

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Almost weekend!")
default:
    fmt.Println("A regular day")
}
```

### 4.3 For Loop

- The only looping structure in Go.
- Can be used as a traditional loop, infinite loop, or range-based loop.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

## 5. Functions

### 5.1 Basic Function

- Function declaration with parameters and return type.

```go
func add(a int, b int) int {
    return a + b
}
```

### 5.2 Multiple Return Values

- Functions can return multiple values, commonly used for error handling.

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### 5.3 Variadic Functions

- Allows a variable number of arguments.

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```

## 6. Structs and Interfaces

### 6.1 Structs

- Group fields together into a single type.

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

### 6.2 Methods

- Define methods on structs for object-oriented programming.

```go
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}
```

### 6.3 Interfaces

- Define a set of method signatures without implementations.
- A type satisfies an interface if it implements all the methods in the interface.

```go
type Greeter interface {
    Greet() string
}

func Introduce(g Greeter) {
    fmt.Println(g.Greet())
}
```

## 7. Pointers

### 7.1 Pointer Declaration and Usage

- Pointers hold the memory address of variables.

```go
func increment(x *int) {
    *x = *x + 1
}

var y = 5
increment(&y)
```

## 8. Error Handling

### 8.1 Error Type and Custom Errors

- Use the `error` type and the `errors.New` or `fmt.Errorf` for custom errors.

```go
func checkValue(val int) error {
    if val <= 0 {
        return errors.New("value must be positive")
    }
    return nil
}
```

### 8.2 Error Checking Pattern

- Errors are checked and handled immediately after calling a function.

```go
if err := checkValue(-1); err != nil {
    fmt.Println("Error:", err)
}
```

## 9. Concurrency

### 9.1 Goroutines

- Lightweight threads managed by Go’s runtime.

```go
go someFunction()
```

### 9.2 Channels

- Channels are used to communicate between goroutines.

```go
ch := make(chan int)
go func() {
    ch <- 5
}()
fmt.Println(<-ch)
```

## 10. Common Idioms and Patterns

### 10.1 Defer

- `defer` schedules a function call to run after the surrounding function completes, commonly used for resource cleanup.

```go
file, _ := os.Open("file.txt")
defer file.Close()
```

### 10.2 Empty Interface (`interface{}`)

- Represents any type, useful for functions needing flexibility with input types.

```go
func printAny(val interface{}) {
    fmt.Println(val)
}
```

### 10.3 Type Assertion

- Used to access the underlying concrete type of an interface.

```go
var i interface{} = "hello"
s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

## 11. Package Management and Project Structure

### 11.1 Packages

- Go code is organized into packages; `main` is the special package for executables.

```go
package mypackage

func MyFunction() {
    // Exported function with capitalized name
}
```

### 11.2 Project Structure

- Go projects follow a directory structure:
  ```
  myproject/
  ├── cmd/            # Main applications
  ├── pkg/            # Library code
  ├── internal/       # Private packages
  ├── main.go         # Entry point for main package
  ```

## 12. Error and Testing

### 12.1 Error Wrapping

- Use `fmt.Errorf` with `%w` for error wrapping, useful in error handling chains.

```go
import "fmt"

err := fmt.Errorf("an error occurred: %w", originalErr)
```

### 12.2 Writing Tests

- Go has a built-in testing library.

```go
package mypackage

import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}
```

---

This cheat sheet covers Go’s essentials and advanced topics, serving as a solid reference to build high-quality, idiomatic Go programs.
