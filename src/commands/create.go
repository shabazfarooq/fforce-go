package commands

import (
	"fmt"
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

  // Set local options
  this.options = options

  // If empty or does not have -t and -n, assume help
  if len(this.options.Options) == 0 {
	  fmt.Println("fforce create [ApexClass, ApexPage, ApexComponent, ApexTrigger] ComponentName")
	  return
	}
  
  // Extract parameters
  this.ComponentType = this.options.Options[0]
  this.ComponentName = this.options.Options[1]

  // Login
  projectio.ExecuteLoginScript(".")

  // Create component
  this.createComponent()

  // Fetch component
  this.fetchComponent()
}

func (this *Create) createComponent() {
	projectio.ExecuteForceShellCommand("create", "-t", this.ComponentType, "-n", this.ComponentName)
}

func (this *Create) fetchComponent() {
	projectio.ExecuteForceShellCommand("fetch", "-t", this.ComponentType, "-n", this.ComponentName)
}