package commands

import (
	"fmt"
	"log"
	"fforce-go/src/projectio"
)

type Build struct {
	options Options
	parentPath string
	buildFilePath string
	buildFileExt string
	buildFileFolderPath string
}

func (this *Build) New(options Options) {
  fmt.Println("** Build **\n")

	// Validate parameters
	this.validateAndSetParams(options)
  
  // // Login
  // projectio.ExecuteLoginScript(this.parentPath)

  // Build
  this.buildFile()
}

func (this *Build) validateAndSetParams(options Options) {
  // Set local options
  this.options = options

  // If options empty or does not have type and name parameters, assume help
  if len(this.options.Options) < 4 {
	  log.Fatal("fforce build [ParentPath] [BuildFilePath] [BuildFileFolderPath] [BuildFileExtension (cls, trigger, page, component, apex, soql, css, html, js, xml)]")
	}

	// Set local properties
  this.parentPath = this.options.Options[0]
  this.buildFilePath = this.options.Options[1]
  this.buildFileFolderPath = this.options.Options[2]
  this.buildFileExt = this.options.Options[3]
}

func (this *Build) buildFile() {
	if this.buildFileExt == "cls" || this.buildFileExt == "trigger" || this.buildFileExt == "page" || this.buildFileExt == "component" {
		// Login
		projectio.ExecuteLoginScript(this.parentPath)

		// Push
		projectio.ExecuteForceShellCommand("push", "-f", this.buildFilePath)

	} else if this.buildFileExt == "apex" {
		// Login
		projectio.ExecuteLoginScript(this.parentPath)

		// Execute anon
		projectio.ExecuteForceShellCommand("apex", this.buildFilePath)

	} else if this.buildFileExt == "soql" {
		// Login
		projectio.ExecuteLoginScript(this.parentPath)

		// Query
		query := projectio.ExtractFirstQueryFromFile(this.buildFilePath)
		fmt.Println("Query: " + query + "\n\n")

		//--format, -f   Output format: csv, json, json-pretty, console
		projectio.ExecuteForceShellCommand("query", "-f", "json", query)
	} else if this.buildFileExt == "css" || this.buildFileExt == "html" || this.buildFileExt == "js" || this.buildFileExt == "xml" {
		projectio.ExecuteShellCommand("sfdx", "force:source:deploy", "-u", "THIS_NEEDS_TO_BE_UPDATED", "-p", this.buildFileFolderPath);
	}
}
