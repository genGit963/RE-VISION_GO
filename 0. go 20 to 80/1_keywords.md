# Go Language Keywords

Go has a set of reserved keywords that serve specific purposes in language syntax. These can be classified into categories based on their usage.

---

## 1. **Package and Import Keywords**

### `package`

Defines the package for a file. Every Go file starts with a `package` declaration.

```go
package main

func main() {
    fmt.Println("Hello, Go!")
}
```

### `import`

Imports other packages for use in the file.

```go
package main

import "fmt"

func main() {
    fmt.Println("Imported the fmt package!")
}
```

---

## 2. **Variable Declaration and Type Keywords**

### `var`

Declares a variable with an optional initializer.

```go
var x int = 10
```

### `const`

Declares a constant value that cannot be modified.

```go
const Pi = 3.14
```

### `type`

Defines a new type, often used with structs, interfaces, or custom types.

```go
type Person struct {
    Name string
    Age  int
}
```

---

## 3. **Function and Method Keywords**

### `func`

Defines a function or method.

```go
func add(x int, y int) int {
    return x + y
}
```

---

## 4. **Control Flow Keywords**

### `if`

Starts a conditional statement.

```go
if x > 10 {
    fmt.Println("x is greater than 10")
}
```

### `else`

Provides an alternative branch in an `if` statement.

```go
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is 10 or less")
}
```

### `for`

Starts a loop, the only loop keyword in Go.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### `switch`

Starts a switch statement, a control structure for multiple conditions.

```go
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Another day")
}
```

### `case`

Defines a branch in a `switch` statement.

```go
switch x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}
```

### `fallthrough`

Allows execution to continue to the next case in a `switch` statement.

```go
switch x {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

### `default`

Specifies the default case in a `switch` statement.

```go
switch x {
default:
    fmt.Println("No case matched")
}
```

---

## 5. **Concurrency Keywords**

### `go`

Starts a goroutine, enabling concurrent execution.

```go
go func() {
    fmt.Println("Running in a goroutine")
}()
```

### `select`

Waits on multiple communication operations, similar to a `switch` for channels.

```go
select {
case msg := <-ch:
    fmt.Println("Received message:", msg)
default:
    fmt.Println("No message received")
}
```

---

## 6. **Channel Keywords**

### `chan`

Declares a channel, used for communication between goroutines.

```go
ch := make(chan int)
```

### `<-`

Used to send or receive data over a channel.

```go
ch <- 10 // Send 10 to channel ch
x := <-ch // Receive from channel ch
```

---

## 7. **Data Structure Keywords**

### `struct`

Defines a struct, a composite data type.

```go
type Point struct {
    X int
    Y int
}
```

### `map`

Declares a map, a key-value data structure.

```go
m := make(map[string]int)
m["age"] = 25
```

### `interface`

Defines an interface, which specifies method signatures.

```go
type Stringer interface {
    String() string
}
```

---

## 8. **Error Handling Keywords**

### `defer`

Defers execution of a function until the surrounding function returns.

```go
defer fmt.Println("This will run last")
```

### `panic`

Triggers a runtime error, terminating the program if not recovered.

```go
panic("Something went wrong!")
```

### `recover`

Used to regain control of a panicking goroutine, allowing cleanup.

```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from", r)
        }
    }()
    panic("Causing a panic!")
}
```

---

## 9. **Pointer and Memory Management Keywords**

### `new`

Allocates memory for a value and returns a pointer to it.

```go
p := new(int)
*p = 42
```

### `make`

Allocates and initializes slices, maps, and channels.

```go
s := make([]int, 5)
```

---

## 10. **Other Control Keywords**

### `return`

Returns from a function, optionally with values.

```go
func sum(a int, b int) int {
    return a + b
}
```

### `break`

Exits a `for` loop or `switch` statement.

```go
for i := 0; i < 5; i++ {
    if i == 3 {
        break
    }
    fmt.Println(i)
}
```

### `continue`

Skips the current iteration of a `for` loop.

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

---

## 11. **Type Keywords**

### `interface`

Defines a set of methods that a type must implement.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### `map`

Defines a map, an associative data type that maps keys to values.

```go
m := make(map[string]int)
m["key"] = 10
```

### `struct`

Defines a structure that groups together multiple fields.

```go
type Employee struct {
    ID   int
    Name string
}
```

---

## 12. **Unused Keywords (for completeness)**

### `goto`

Directly transfers control to a labeled statement. Rarely used in modern Go.

```go
func example() {
    i := 0
Label:
    fmt.Println(i)
    i++
    if i < 3 {
        goto Label
    }
}
```

### `select`

Used with channels to handle multiple channel operations.

```go
select {
case msg := <-ch1:
    fmt.Println("Received from channel 1:", msg)
case msg := <-ch2:
    fmt.Println("Received from channel 2:", msg)
}
```

---

Each keyword in Go is specifically designed to make the language simple yet powerful, balancing modern features with easy-to-read syntax. This list provides a strong foundation to understand Go’s functional and idiomatic structure.
