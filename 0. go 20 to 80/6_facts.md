# FACTs about go-lang

## 1. **Keywords** in Go

- **Total Keywords**: Go has only 25 reserved keywords, making it minimalistic and easy to remember. Examples include `func`, `var`, `const`, `package`, `import`, `defer`, and `select`.
- **No Generic Keyword**: Go initially launched without generics (introduced only in Go 1.18), relying instead on interfaces and concrete types.
- **No `class` or `inheritance` Keywords**: Go avoids traditional inheritance and class hierarchies, opting for composition and interfaces instead.
- **Case Sensitivity**: Go distinguishes between uppercase and lowercase, which affects whether identifiers (functions, variables) are exported or kept private.

## 2. **Symbols and Operators**

- **:= Operator**: This is unique to Go and allows concise variable declaration and initialization (e.g., `x := 10`).
- **`&` and `*` for Pointers**: The `&` symbol gets a variable’s address, while `*` dereferences a pointer, allowing access to the value at that memory location.
- **Only `for` for Loops**: Go has only one looping keyword, `for`, making it versatile enough for traditional, infinite, and range-based loops.
- **Range in For Loops**: The `range` keyword simplifies iterating over arrays, slices, maps, and channels.

## 3. **Rules and Structure**

- **No Semicolons**: Go uses implicit semicolon insertion, so you generally don’t need to end statements with semicolons.
- **No Ternary Operator**: Go avoids the ternary conditional operator (`condition ? true : false`) to keep the language simple and readable.
- **Exported Names**: In Go, only identifiers (variables, functions, etc.) starting with an uppercase letter are accessible outside their package.
- **No Circular Dependencies**: Go prevents circular dependencies across packages, promoting a cleaner architecture.
- **Error Handling via Multiple Returns**: Unlike exceptions in many languages, Go handles errors by returning both a value and an error (e.g., `value, err := someFunction()`).

## 4. **Functions and Methods**

- **Multiple Return Values**: Go functions can return multiple values, which is often used for value-error pairs.
- **Defer**: Go’s `defer` statement is unique and schedules a function to run after the current function completes, useful for resource cleanup.
- **Variadic Parameters**: Functions can accept a variable number of arguments of the same type using `...`, like `func sum(nums ...int) int`.

## 5. **Types and Interfaces**

- **No Implicit Type Conversion**: Go requires explicit conversions between types (e.g., `int(x)`), avoiding implicit coercion.
- **Interface Satisfaction**: Go’s interfaces are satisfied implicitly. If a type implements all the methods of an interface, it automatically satisfies that interface without needing any special keyword or declaration.
- **Empty Interface**: `interface{}` is Go’s way to represent any type, as it doesn’t restrict types that can be assigned to it.

## 6. **Concurrency**

- **Goroutines**: Go introduced `goroutines`, lightweight threads managed by the Go runtime. Starting a goroutine is as simple as prefixing a function call with `go`.
- **Channels**: Go uses `channels` for safe communication between goroutines, which prevents data races and is central to its concurrency model.
- **Select Statement**: Go has a `select` statement for handling multiple channel operations simultaneously, similar to a `switch` but for channels.

## 7. **Memory Management**

- **Garbage Collection**: Go has built-in garbage collection, which automatically manages memory allocation and deallocation.
- **No Manual Memory Management**: Go abstracts memory management entirely away from the developer, simplifying code and reducing bugs related to memory.

## 8. **Built-in Testing**

- **Go’s Built-in Testing**: Go has a simple testing framework included in the language itself (`testing` package), promoting testing as a standard part of Go development.

## 9. **Project Structure**

- **Standard Project Structure**: Go has a recommended structure (`cmd/`, `pkg/`, `internal/`) that helps developers organize large projects, making code modular and maintainable.
- **No Hidden Imports**: Go’s `import` statement is explicit—each imported package must be used, and unused imports cause compilation errors, keeping code clean.

## 10. **Code Style and Conventions**

- **Gofmt for Formatting**: Go includes `gofmt`, a built-in tool to enforce a consistent style across Go codebases, saving time and avoiding style-related debates.
- **Error Wrapping**: Go encourages error wrapping with `%w` for error handling chains, preserving the original error context while adding details.
- **Avoiding Nulls**: Go prefers `nil` over null and avoids nullable values wherever possible, reducing risks of null reference errors.

---

These elements define Go’s simplicity, efficiency, and uniqueness.
