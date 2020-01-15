package commands

import (
	"fmt"
	"../projectio"
)

type Build struct {
	options Options
	parentPath string
	buildFilePath string
	buildFileExt string
}

func (this *Build) New(options Options) {
  fmt.Println("** Build **\n")

  // Set local options
  this.options = options

  // TODO: validate parameters
  
  // Extract parameters
  this.parentPath = this.options.Options[0]
  this.buildFilePath = this.options.Options[1]
  this.buildFileExt = this.options.Options[2]

  // Login
  projectio.ExecuteLoginScript(this.parentPath)

  // Build
  this.buildFile()
}

func (this *Build) buildFile() {
	if this.buildFileExt == "cls" || this.buildFileExt == "trigger" || this.buildFileExt == "page" || this.buildFileExt == "component" {
		// Push
		projectio.ExecuteForceShellCommand("push", "-f", this.buildFilePath)

	} else if this.buildFileExt == "apex" {
		// Execute anon
		projectio.ExecuteForceShellCommand("apex", this.buildFilePath)

	} else if this.buildFileExt == "soql" {
		// Query
		query := projectio.ExtractFirstQueryFromFile(this.buildFilePath)
		fmt.Println("Query: " + query + "\n\n")

		projectio.ExecuteForceShellCommand("query", "-f", "console", query)
	}
}