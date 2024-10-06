package main

import (
	"embed"
	"fmt"
)

/*
---> Embed directive <---
"//go:embed" is a compiler directive
 that allows programs to include arbitrary files and folders
 in the Go binary at build time.

 Read more about the embed directive here.
*/

func learn_embed_directive() {
	fmt.Println("\n------------ learn_embed_directive --------------")

	/*
		embed directives accept paths relative to the directory
		containing the Go source file.
		This directive embeds the contents of the file
		into the string variable immediately following it.
	*/

	// "//go:embed" folder/single_file.txt
	var fileString string

	// Or embed the contents of the file into a []byte.
	// "//go:embed" folder/single_file.txt
	var fileByte []byte

	/*
		We can also embed multiple files or even folders with wildcards.
		This uses a variable of the embed.FS type,
		which implements a simple virtual file system.
	*/
	// "//go:embed" folder/single_file.txt
	// "//go:embed" folder/*.hash
	var folder embed.FS

	// Print out the contents of single_file.txt.
	print(fileString)
	print(string(fileByte))

	// Retrieve some files from the embedded folder.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

// func main() {
// 	learn_embed_directive()
// 	println("\n-------------------------------")
// }
