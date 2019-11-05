package commands

import (
  "fmt"
  "log"
  "../projectio"
  "../sfdcapi"
)

type Init struct {
  username string
  password string
  instanceUrl string
  instanceType string
  sessionId string
}

func (this *Init) New(options []string) {
  fmt.Println("Executing Init")
  // fmt.Println(options)
  // hasOption()
  log.Fatal("leaving..")

  // Capture user credentials
  this.askUserForCredentials()
  
  // Authenticate with SFDC
  authenticatedCredentials := sfdcapi.AuthenticateToSFDC(
    this.username,
    this.password,
    this.instanceUrl,
  )

  // Update instanceUrl with verified baseUrl from SFDC
  this.instanceUrl = authenticatedCredentials.BaseUrl

  // Create local files if credentials authenticated
  if authenticatedCredentials.Authenticated {
    fmt.Println("\n** Authentication Successful **\n");
    this.createLocalFilesAndDirectories();
  } else {
    fmt.Println("\n** Authentication Failed **\n");
  }
}

/**
 * Capture user credentials
 */
func (this *Init) askUserForCredentials() {
  // Capture user credentials
  this.username = projectio.AskUser("Enter username")
  this.password = this.askPassword()
  this.instanceType = projectio.AskUser("Enter instance type(test/login/full URL)")
 
  // Finalize instance URL
  if this.instanceType == "test" || this.instanceType == "login" {
    this.instanceUrl = "https://" + this.instanceType + ".salesforce.com";
  } else {
    this.instanceUrl = this.instanceType
  }

  // Validate user credentials
  if this.username == "" || this.password == "" || this.instanceType == "" {
    log.Fatal("Missing username/password/instance type")
  }
}

func (this *Init) askPassword() string {
  // const hidePassword = super.hasOption('showpassword') === false;
  hidePassword := false;

  if hidePassword {
    return projectio.AskUserPassword("Enter password")
  } else {
    return projectio.AskUser("Enter password")
  }
}

/**
 * Create local files and directories
 */
func (this *Init) createLocalFilesAndDirectories() {
  projectio.CreateOpenUrlFile(
    this.username,
    this.password,
    this.instanceUrl,
  )
  projectio.CreateBuildPropertiesFile(
    this.username,
    this.password,
    this.instanceUrl,
  )
  projectio.CreateLoginFile(
    this.username,
    this.password,
    this.instanceType,
  )
  projectio.CreateBuildXmlFile()
  projectio.CreatePackageXml();
  projectio.CreateExecuteAnonFile();
  projectio.CreateQueryFile();
}