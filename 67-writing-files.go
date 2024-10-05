package main

import (
	"fmt"
	"os"
)

/*
---> Writing File <---
*/

func checking(e error) {
	if e != nil {
		panic(e)
	}

}
func learn_writing_file() {
	fmt.Println("\n------------ learn_writing_file --------------")

	// To start, here’s how to dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("sample_write1.txt", d1, 0644)
	checking(err)

	// For more granular writes, open a file for writing.
	f, err := os.Create("sample_write2.txt")
	checking(err)

	// It’s idiomatic to "defer" a Close immediately after opening a file.
	defer f.Close()

	// You can Write byte slices as you’d expect.
	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, err := f.Write(d2)
	// checking(err)
	// fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.

	//     n3, err := f.WriteString("writes\n")
	//     checking(err)
	//     fmt.Printf("wrote %d bytes\n", n3)
	// Issue a Sync to flush writes to stable storage.

	//     f.Sync()
	// bufio provides buffered writers in addition to the buffered readers we saw earlier.

	//     w := bufio.NewWriter(f)
	//     n4, err := w.WriteString("buffered\n")
	//     checking(err)
	//     fmt.Printf("wrote %d bytes\n", n4)
	// Use Flush to ensure all buffered operations have been applied to the underlying writer.

	// w.Flush()
}

func main() {
	learn_writing_file()
	println("\n-------------------------------")
}
