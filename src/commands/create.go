package commands

import (
	"fmt"
	"os/exec"
	// "bytes"
	"log"
	"strings"
	// "io"
	// "sync"
	// "../projectio"
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
  loginSucessful := this.login()
  if !loginSucessful {
  	log.Fatal("Unable to login with force")
  }

  // Create component
  this.createComponent()

  // Fetch component
  this.fetchComponent()
}

// TODO: Dont login unless necessary! save time
func (this *Create) login() bool {
	loginCommand := "./login"
	result := executeShellCommand(loginCommand)

	return strings.Contains(result, "Logged in as")
}

func (this *Create) createComponent() {
	cmd := exec.Command("force", "create", "-t", this.ComponentType, "-n", this.ComponentName)
	executeForceShellCommand(cmd)
}

func (this *Create) fetchComponent() {
	cmd := exec.Command("force", "fetch", "-t", this.ComponentType, "-n", this.ComponentName)
	executeForceShellCommand(cmd)
}