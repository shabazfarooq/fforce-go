package projectio

import (
	"fmt"
	"os/exec"
	"bytes"
	"log"
	"io"
	"sync"
	"strings"
)

/**
 * Public shell commands
 */

// TODO: Dont login unless necessary! save time
func ExecuteLoginScript(parentPath string) bool {
	result := executeShellCommand(parentPath + "/login")

	loginSucessful := strings.Contains(result, "Logged in as")
	if !loginSucessful {
		log.Fatal("Unable to login with force")
	}

	return loginSucessful
}

func ExecuteForceShellCommand(commandArgs ...string) {
	cmd := exec.Command("force", commandArgs...)
	executeShellCommand2(cmd)
}

/**
 * Private shell commands
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

func executeShellCommand2(cmd *exec.Cmd) {
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
