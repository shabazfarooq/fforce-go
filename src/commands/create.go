package commands

import (
	"fmt"
	"log"
	"../projectio"
)

type Create struct {
	options Options
	ComponentType string
	ComponentName string
	CreateFileExt string
}

func (this *Create) New(options Options) {
  fmt.Println("** Create **\n")

  // Validate parameters
  this.validateAndSetParams(options)

  // Login
  projectio.ExecuteLoginScript(".")

  // Create component
  this.createComponent()

  // Fetch component
  this.fetchComponent()
}

func (this *Create) validateAndSetParams(options Options) {
  // Set local options
  this.options = options

  // If options empty or does not have type and name parameters, assume help
  if len(this.options.Options) < 2 {
	  log.Fatal("fforce create [ApexClass, ApexPage, ApexComponent, ApexTrigger] [ComponentName]")
	}

	// Set local properties
  this.ComponentType = this.options.Options[0]
  this.ComponentName = this.options.Options[1]
}

func (this *Create) createComponent() {
	projectio.ExecuteForceShellCommand("create", "-t", this.ComponentType, "-n", this.ComponentName)
}

func (this *Create) fetchComponent() {
	projectio.ExecuteForceShellCommand("fetch", "-t", this.ComponentType, "-n", this.ComponentName)
}