package cli

import (
	"fmt"
	"os"
)

// Prints a message and exits the program.
func die(format string, a ...interface{}) {
	formatStr := fmt.Sprintf("ERROR: %s\n", format)
	fmt.Fprintf(os.Stderr, formatStr, a...)
	os.Exit(1)
}
