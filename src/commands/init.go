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
  securityToken string
  instanceUrl string
  instanceType string
  sessionId string
  options Options
}

func (this *Init) New(options Options) {
  fmt.Println("** Init **\n")

  // Set local options
  this.options = options

  // Capture user credentials
  this.askUserForCredentials()
  
  // Authenticate with SFDC
  authenticatedCredentials := sfdcapi.AuthenticateToSFDC(
    this.username,
    this.password,
    this.securityToken,
    this.instanceUrl,
  )

  // Create local files if credentials authenticated
  if authenticatedCredentials.Authenticated {
    fmt.Println("\n** Authentication Successful **\n");
    // Update instanceUrl with verified baseUrl from SFDC
    this.instanceUrl = authenticatedCredentials.BaseUrl
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
  this.securityToken = projectio.AskUser("Enter security token (optional)")
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
  hidePassword := this.options.hasOption("h", "hidePassword")

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
    this.securityToken,
    this.instanceUrl,
    this.instanceType,
  )
  projectio.CreateLoginFile(
    this.username,
    this.password,
    this.securityToken,
    this.instanceType,
  )
  projectio.CreateBuildXmlFile()
  projectio.CreatePackageXml();
  projectio.CreateExecuteAnonFile();
  projectio.CreateQueryFile();
}
