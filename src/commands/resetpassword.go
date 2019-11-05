package commands

import (
  "fmt"
  "log"
  "../projectio"
  "../sfdcapi"
)

type ResetPassword struct {
  username string
  password string
  instanceUrl string
  instanceType string
  options Options
}

func (this *ResetPassword) New(options Options) {
  // Set local options
  this.options = options

  fmt.Println("RESET PASSWORD")

  
  // Determine existing credentials
  this.determineExistingCredentials()



  log.Fatal("BYE")



  
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
func (this *ResetPassword) determineExistingCredentials() {
  res := projectio.ReadBuildProperties()
  this.username = res.Username
  this.instanceUrl = res.ServerUrl
  this.instanceType = res.InstanceType



  // Validate user credentials
  if this.username == "" || this.instanceUrl == "" || this.instanceType == "" {
    log.Fatal("Missing username/password/instance type in build.properties file")
  }
}

func (this *ResetPassword) askPassword() string {
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
func (this *ResetPassword) createLocalFilesAndDirectories() {
  projectio.CreateOpenUrlFile(
    this.username,
    this.password,
    this.instanceUrl,
  )
  projectio.CreateBuildPropertiesFile(
    this.username,
    this.password,
    this.instanceUrl,
    "blah",
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