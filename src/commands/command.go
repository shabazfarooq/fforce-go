package commands

/**
 * Command Interface
 */
type Command interface {
  New(options Options)
}