package commands

import (
  "fmt"
  "log"
  "fforce-go/src/projectio"
  "fforce-go/src/sfdcapi"
)

type ResetPassword struct {
  username string
  password string
  securityToken string
  instanceUrl string
  instanceType string
  options Options
}

func (this *ResetPassword) New(options Options) {
  fmt.Println("** Reset Password **\n")

  // Set local options
  this.options = options
  
  // Determine existing credentials
  this.determineExistingCredentials()

  // Get new password
  this.password = this.askPassword()

  // Get new security token
  this.securityToken = projectio.AskUser("Enter security token (optional)")

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
    log.Fatal("\n** Authentication Failed **\n");
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
}
