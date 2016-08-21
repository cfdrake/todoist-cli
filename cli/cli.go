package cli

import (
  "fmt"
  "os"
  "os/user"
)

// Runs the command line interface with the given argument list.
func Execute(args []string) {
  // Fetch the current user's config.
  me, err := user.Current()
  if err != nil {
    fmt.Fprintln(os.Stderr, "ERROR: Could not find current user!")
    os.Exit(1)
  }

  _, err = loadConfiguration(me)
  if err != nil {
    fmt.Fprintln(os.Stderr, "ERROR:", err)
    os.Exit(1)
  }
}
