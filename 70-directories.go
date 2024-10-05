package main

import (
	"fmt"
	"os"
)

/*
---> Directories <---
Go has several useful functions for working with directories in the file system.
*/

func checking_error(e error) {
	if e != nil {
		panic(e)
	}
}

func learn_directories() {
	fmt.Println("\n------------ learn_directories --------------")

	// Create a new sub-directory in the current working directory.
	err := os.Mkdir("subdir", 0755)
	checking_error(err)

	/*
		When creating temporary directories,
		it’s good practice to defer their removal.
		os.RemoveAll will delete a whole directory tree (similarly to rm -rf).
	*/
	defer os.RemoveAll("subdir")

	// Helper function to create a new empty file.
	createEmptyFile := func(name string) {
		d := []byte("")
		checking_error(os.WriteFile(name, d, 0644))
	}
	createEmptyFile("subdir/file1")
	/*
		We can create a hierarchy of directories,
		 including parents with MkdirAll. This is similar to the command-line mkdir -p.
	*/
	err = os.MkdirAll("subdir/parent/child", 0755)
	checking_error(err)
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir lists directory contents, returning a slice of os.DirEntry objects.
	//     c, err := os.ReadDir("subdir/parent")
	//     checking_error(err)
	//     fmt.Println("Listing subdir/parent")
	//     for _, entry := range c {
	//         fmt.Println(" ", entry.Name(), entry.IsDir())
	//     }
	// Chdir lets us change the current working directory, similarly to cd.

	//     err = os.Chdir("subdir/parent/child")
	//     checking_error(err)
	// Now we’ll see the contents of subdir/parent/child when listing the current directory.

	//     c, err = os.ReadDir(".")
	//     checking_error(err)
	//     fmt.Println("Listing subdir/parent/child")
	//     for _, entry := range c {
	//         fmt.Println(" ", entry.Name(), entry.IsDir())
	//     }
	// cd back to where we started.

	//     err = os.Chdir("../../..")
	//     checking_error(err)
	// We can also visit a directory recursively, including all its sub-directories. WalkDir accepts a callback function to handle every file or directory visited.

	// fmt.Println("Visiting subdir")
	// err = filepath.WalkDir("subdir", visit)
}

func main() {
	learn_directories()
	println("\n-------------------------------")
}
