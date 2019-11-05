package commands

import "fmt"

/**
 * Command Interface
 */
type Command interface {
  New(option string)
}

/**
 * Command Package Methods
 */
func hasOption(option string) bool {
  fmt.Println("in super hasoption method")
  return true
}

