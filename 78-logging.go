package main

import (
	"fmt"
	"log"
)

/*
---> Logging <---
The Go standard library provides straightforward tools for outputting logs
from Go programs,

with the log package for free-form output
and the log/slog package for structured output.
*/

func learn_logging() {
	fmt.Println("\n------------ learn_logging --------------")

	/*
		Simply invoking functions like Println from the log package
		uses the standard logger,

		which is already pre-configured for reasonable
		logging output to os.Stderr.
		Additional methods like Fatal* or Panic* will exit the program after logging.
	*/
	log.Println("standard logger")

	/*
		Loggers can be configured with flags to set their output format.

		By default, the standard logger has the log.Ldate
		and log.Ltime flags set, and these are collected in log.LstdFlags.

		We can change its flags to emit time with microsecond accuracy,

		for example.
	*/
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

}

func main() {
	learn_logging()
	println("\n-------------------------------")
}
