package projectio

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "bufio"
  "strings"
  "golang.org/x/crypto/ssh/terminal"
  "syscall"
)

/**
 * User input
 */
func AskUser(question string) string {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print(question + ": ")
  text, err := reader.ReadString('\n')

  if err != nil {
    log.Fatal(err)
  }

  return strings.TrimSuffix(text, "\n")
}

func AskUserPassword(question string) string {
  fmt.Print(question + ": ")
  bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))

  if err != nil {
    log.Fatal(err)
  }

  password := string(bytePassword)
  fmt.Println()
  
  return password
}

/**
 * Read method
 */
func readFile(filename string) string {
  content, err := ioutil.ReadFile(filename)
  
  if err != nil {
    log.Fatal(err)
  }

  return string(content);
}

/**
 * Write method
 */
func writeToFile(filename string, textToWrite string) {
  file, err := os.Create(filename)
  
  if err != nil {
    log.Fatal("Cannot create file: ", err)
  }

  defer file.Close()

  fmt.Fprintf(file, textToWrite)
}

/**
 * CHMOD method
 */
func makeFileExecutable(filename string) {
  // err := os.Chmod(filename, 777)

  cmd := exec.Command("chmod", "+x", filename)
  _, err := cmd.Output()

  if err != nil {
    log.Fatal("Cannot chmod 777 file: ", err);
  }
}

/**
 * Make directory method
 */
func makeDirectory(directoryName string) {
  _ = os.Mkdir(directoryName, 0755)
}
