package cli

import (
	"fmt"
	"os"
)

// Prints a message and exits the program.
func die(format string, a ...interface{}) {
	fmt.Fprint(os.Stderr, "ERROR: ")
	fmt.Fprintf(os.Stderr, format, a...)
	fmt.Println()
	os.Exit(1)
}
