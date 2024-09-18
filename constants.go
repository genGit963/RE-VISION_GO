package main

import (
	"fmt"
	"math"
)

// global contant
const s string = "i am constant"

func learnAboutConstant() {
	fmt.Println("global constant s: ", s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("d: ", d)

	fmt.Println("integer bit 64 converter: ", int64(d))

	fmt.Println("sin(n): ", math.Sin(n))
}

func main() {
	learnAboutConstant()
}
