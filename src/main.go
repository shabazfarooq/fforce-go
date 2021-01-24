package main

import (
	"fmt"
	"os"
	"log"
	"../src/commands"
)

type BeginProperties struct {
  commandRequested string
  options []string
}

func main() {
	// Greet
	greet()

	// Get begin properties
	beginProperties := getBeginProperties()

	// Execute command
	executeRequestedCommand(beginProperties)
}


/**
 * Helper methods
 */
func greet() {
	fmt.Println("** FForce **")
}

func getBeginProperties() BeginProperties {
	var argsLength = len(os.Args)
	var returnStruct BeginProperties

	// Extract command (index = 1)
	if argsLength >= 2 {
		returnStruct.commandRequested = os.Args[1]
	} else {
		log.Fatal("Missing command, ie: init, build, create, reset-password")
	}

	// Extract options (index >=2)
	if argsLength >= 3 {
		returnStruct.options = os.Args[2:]
	} else {
		returnStruct.options = []string{}
	}

	return returnStruct
}

func executeRequestedCommand(beginProperties BeginProperties) {
	// Build options type
	options := commands.Options{beginProperties.options}

	// Determine command to execute
	var command commands.Command

	if beginProperties.commandRequested == "init" {
		command = &commands.Init{}
	} else if beginProperties.commandRequested == "reset-password" {
		command = &commands.ResetPassword{}
	} else if beginProperties.commandRequested == "build" {
		command = &commands.Build{}
	} else if beginProperties.commandRequested == "create" {
		command = &commands.Create{}
	} else {
		log.Fatal("Unknown command: '" + beginProperties.commandRequested + "'. try: init, build, create, reset-password")
	}

	// Execute command with options
	command.New(options)
}
