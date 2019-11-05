package commands

import (
	"fmt"
	"os/exec"
	"bytes"
	"log"
	"strings"
	"io"
	"sync"
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
  loginSucessful := this.login()
  if !loginSucessful {
  	log.Fatal("Unable to login with force")
  }

  // Build
  this.buildFile()
}

// TODO: Dont login unless necessary! save time
func (this *Build) login() bool {
	loginCommand := this.parentPath + "/login"
	result := executeShellCommand(loginCommand)

	return strings.Contains(result, "Logged in as")
}

func (this *Build) buildFile() {
	if this.buildFileExt == "cls" || this.buildFileExt == "trigger" || this.buildFileExt == "page" || this.buildFileExt == "component" {
		// Push
		
		cmd := exec.Command("force", "push", "-f", this.buildFilePath)
		executeForceShellCommand(cmd)

	} else if this.buildFileExt == "apex" {
		// Execute anon
		
		cmd := exec.Command("force", "apex", this.buildFilePath)
		executeForceShellCommand(cmd)

	} else if this.buildFileExt == "soql" {
		// Query

		query := projectio.ExtractFirstQueryFromFile(this.buildFilePath)
		fmt.Println("Query: " + query + "\n\n")

		cmd := exec.Command("force", "query", "-f", "console", query)
		executeForceShellCommand(cmd)
	}
}





/**
 * Shell stuff below
 */





// TODO: Error handling... this just breaks when it breaks
func executeShellCommand(command string) string {
	cmd := exec.Command(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	
	result := string(out.Bytes())
	fmt.Println(result)

	return result;
}

func executeForceShellCommand(cmd *exec.Cmd) {
	var wg sync.WaitGroup

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		copyLogs(stdout)
	}()

	go func() {
		defer wg.Done()
		copyLogs(stderr)
	}()

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
	}
}

func copyLogs(r io.Reader) {
	buf := make([]byte, 100)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Println(string(buf[0:n]))
		}
		if err != nil {
			break
		}
	}
}
