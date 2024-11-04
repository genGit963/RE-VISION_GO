# Go Language Symbols and Their Meanings

This document provides a comprehensive list of symbols used in the Go programming language, organized by functional categories.

## 1. **Arithmetic Operators**

These operators perform basic arithmetic operations.

| Symbol | Meaning             | Example |
| ------ | ------------------- | ------- |
| `+`    | Addition            | `a + b` |
| `-`    | Subtraction         | `a - b` |
| `*`    | Multiplication      | `a * b` |
| `/`    | Division            | `a / b` |
| `%`    | Modulus (Remainder) | `a % b` |

## 2. **Relational Operators**

These operators compare two values and return a boolean result.

| Symbol | Meaning                  | Example  |
| ------ | ------------------------ | -------- |
| `==`   | Equal to                 | `a == b` |
| `!=`   | Not equal to             | `a != b` |
| `>`    | Greater than             | `a > b`  |
| `<`    | Less than                | `a < b`  |
| `>=`   | Greater than or equal to | `a >= b` |
| `<=`   | Less than or equal to    | `a <= b` |

### 3. Logical Operators

These operators perform logical operations and return boolean values.

- `&&`: Logical AND
  - Example: `a && b`
- `||`: Logical OR
  - Example: `a || b`
- `!`: Logical NOT
  - Example: `!a`

---

### 4. Bitwise Operators

These operators perform bit-level operations on integers.

- `&`: Bitwise AND
  - Example: `a & b`
- `|`: Bitwise OR
  - Example: `a | b`
- `^`: Bitwise XOR
  - Example: `a ^ b`
- `&^`: Bit clear (AND NOT)
  - Example: `a &^ b`
- `<<`: Left shift
  - Example: `a << b`
- `>>`: Right shift
  - Example: `a >> b`

---

### 5. Assignment Operators

These operators assign values to variables.

- `=`: Assignment
  - Example: `a = 5`
- `+=`: Add and assign
  - Example: `a += b`
- `-=`: Subtract and assign
  - Example: `a -= b`
- `*=`: Multiply and assign
  - Example: `a *= b`
- `/=`: Divide and assign
  - Example: `a /= b`
- `%=`: Modulus and assign
  - Example: `a %= b`
- `&=`: Bitwise AND and assign
  - Example: `a &= b`
- `|=`: Bitwise OR and assign
  - Example: `a |= b`
- `^=`: Bitwise XOR and assign
  - Example: `a ^= b`
- `<<=`: Left shift and assign
  - Example: `a <<= b`
- `>>=`: Right shift and assign
  - Example: `a >>= b`

## 6. **Increment and Decrement Operators**

These operators increment or decrement the value of a variable.

| Symbol | Meaning   | Example |
| ------ | --------- | ------- |
| `++`   | Increment | `a++`   |
| `--`   | Decrement | `a--`   |

## 7. **Miscellaneous Operators**

These operators serve special purposes in Go.

| Symbol | Meaning                                   | Example               |
| ------ | ----------------------------------------- | --------------------- |
| `:=`   | Short variable declaration                | `a := 5`              |
| `...`  | Variadic function argument                | `func f(args ...int)` |
| `@`    | Decorator syntax (for specific use cases) | N/A                   |

## 8. **Grouping and Control Flow Symbols**

These symbols help control the flow of the program.

| Symbol | Meaning             | Example                           |
| ------ | ------------------- | --------------------------------- |
| `(`    | Opening parenthesis | `if (a > b) {...}`                |
| `)`    | Closing parenthesis | `for i := 0; i < 10; i++ { ... }` |
| `{`    | Opening brace       | `func main() { ... }`             |
| `}`    | Closing brace       | `}`                               |
| `[`    | Opening bracket     | `arr := []int{1, 2, 3}`           |
| `]`    | Closing bracket     | `}`                               |
| `,`    | Comma               | `func add(a, b int) { ... }`      |
| `;`    | Semicolon           | `if a > b; { ... }`               |

## Conclusion

Understanding these symbols is crucial for writing effective Go programs. They are fundamental to operations, comparisons, and control flow in your applications.
