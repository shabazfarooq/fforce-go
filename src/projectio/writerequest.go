package projectio

import (
  "fmt"
  "strconv"
)

func CreateBuildPropertiesFile(username string,
                               password string,
                               serverurl string,
                               instancetype string) {
  textToWrite := buildProperties(username, password, serverurl, instancetype)
  writeToFile(BUILDPROPERTIES, textToWrite);
  fmt.Println("... Created build.properties file");
}

func CreateBuildXmlFile() {
  textToWrite := buildXml()
  writeToFile(BUILDXML, textToWrite)
  fmt.Println("... Created build.xml file");
}

func CreatePackageXml() {
  makeDirectory(SRCFOLDER)
  fmt.Println("... Created " + SRCFOLDER + " folder");

  fileName := SRCFOLDER + "/package.xml"
  textToWrite := packageXml()
  writeToFile(fileName, textToWrite)
  fmt.Println("... Created " + fileName + " file");
}

func CreateExecuteAnonFile() {
  makeDirectory(EXECUTEANONFOLDER)
  fmt.Println("... Created " + EXECUTEANONFOLDER + " folder")

  for i := 0; i < 3; i++ {
    iStr := strconv.Itoa((i+1))
    fileName := EXECUTEANONFOLDER + "/executeAnon" + iStr + ".apex"
    textToWrite := executeAnon(iStr);
    writeToFile(fileName, textToWrite)
    fmt.Println("... Created " + fileName + " file");
  }
}

func CreateQueryFile() {
  makeDirectory(QUERYFOLDER)
  writeToFile(QUERYFOLDER + "/query.soql", "[SELECT Id FROM Account LIMIT 1]")
  fmt.Println("... Created " + QUERYFOLDER + " folder");
}

func CreateOpenUrlFile(username string, password string, instanceUrl string) {
  textToWrite := openUrl(username, password, instanceUrl);
  filename := "openUrl"
  writeToFile(filename, textToWrite)
  makeFileExecutable(filename)
  fmt.Println("... Created openUrl file");
}

func CreateLoginFile(username string, password string, instanceType string) {
  textToWrite := login(username, password, instanceType);
  filename := "login"
  writeToFile(filename, textToWrite)
  makeFileExecutable(filename)
  fmt.Println("... Created login file");
}