package cli

import (
	"fmt"
	"os"
)

// Prints a message and exits the program.
func die(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a)
	os.Exit(1)
}
