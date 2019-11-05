package main

/****
  ******* 1) implement options
  ******* 2) implement query command
             a) takes in param for query file location
             b) read content
             c) auth to sfdc
						 d) pass in contents
 - This should be able to run by passing in a file, and having it auto detect what should be happening? and passing args?

-- Add chmod capabilities for openUrl
-- Create ./login file as well
-- in init function, use the url from authCreds for creating openUrl file


  implement logger - replace all log.Fatal with an error method in projectio/logger
  finish implementing init function (missing the SFDCAPI stuff)
  change "src" var to be src and not src2 in projectio/constants

  change command interface to use inheritance

  // Command to implement. Open salesforce URL and login
*/

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
	fmt.Println("Running...")
	fmt.Println("")
	fmt.Println("")
}

func getBeginProperties() BeginProperties {
	var argsLength = len(os.Args)
	var returnStruct BeginProperties

	// Extract command (index = 1)
	if argsLength >= 2 {
		returnStruct.commandRequested = os.Args[1]
	} else {
		log.Fatal("Missing command, ie: init, build, reset-password")
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
	var command commands.Command

	// Determine command to execute
	if beginProperties.commandRequested == "init" {
		command = &commands.Init{}
	} else {
		log.Fatal("Unknown command: '" + beginProperties.commandRequested + "'. try: init, build, reset-password")
	}

	// Execute command with options
	command.New(beginProperties.options)
}