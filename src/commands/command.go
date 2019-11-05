package commands

import "fmt"

/**
 * Command Interface
 */
type Command interface {
  New(options []string)
}

/**
 * Command Package Methods
 */
// func hasOption(options []string) bool {
func hasOption() bool {
  fmt.Println("in super hasoption method:")
  // fmt.Println(options)
  return true
}

