package cli

import (
	"fmt"
	"os"
)

// Prints a message and exits the program.
func die(a ...interface{}) {
	fmt.Fprint(os.Stderr, "ERROR: ")
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}
