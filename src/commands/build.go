package commands

import (
	"fmt"
	"log"
	"strings"
	"fforce-go/src/projectio"
)

type Build struct {
	options Options
	parentPath string
	buildFilePath string
	buildFileExt string
	buildFileFolderPath string
	buildFileParentPath string
}

func (this *Build) New(options Options) {
  fmt.Println("** Build **\n")

	// Validate parameters
	this.validateAndSetParams(options)
  
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

  buildFilePathSplit := strings.Split(this.buildFileFolderPath, "/")
  this.buildFileParentPath = buildFilePathSplit[len(buildFilePathSplit) - 2]
}

func (this *Build) buildFile() {
	//
	//
	// SFDX Route
	//
	//
	if this.buildFileExt == "css" || this.buildFileExt == "html" || this.buildFileExt == "js" || this.buildFileExt == "xml" {
		projectio.ExecuteShellCommand("sfdx", "force:source:deploy", "-u", "THIS_NEEDS_TO_BE_UPDATED", "-p", this.buildFileFolderPath);
		return
	}
	
	//
	//
	// Force Route
	//
	//

	// Login
	projectio.ExecuteLoginScript(this.parentPath)


	if this.buildFileParentPath == "aura" {
		projectio.ExecuteForceShellCommand("push", this.buildFileFolderPath)

	} else if this.buildFileExt == "cls" || this.buildFileExt == "trigger" || this.buildFileExt == "page" || this.buildFileExt == "component" {
		projectio.ExecuteForceShellCommand("push", "-f", this.buildFilePath)

	} else if this.buildFileExt == "apex" {
		projectio.ExecuteForceShellCommand("apex", this.buildFilePath)

	} else if this.buildFileExt == "soql" {
		// Query
		query := projectio.ExtractFirstQueryFromFile(this.buildFilePath)
		fmt.Println("Query: " + query + "\n\n")

		//--format, -f   Output format: csv, json, json-pretty, console
		projectio.ExecuteForceShellCommand("query", "-f", "json", query)
	}
}
