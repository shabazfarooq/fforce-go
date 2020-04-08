package projectio

import (
  "strings"
  "log"
)

type BuildPropertiesReadResult struct {
  Username string
  Password string
  SecurityToken string
  ServerUrl string
  InstanceType string
}

func ExtractFirstQueryFromFile(queryFilePath string) string {
  // Get file as string
  fileContents := readFile(queryFilePath)
  
  if (fileContents == "") {
    log.Fatal("Empty file, No query found in file, ensure the query is surrounded by square brackets. ie:\n[SELECT Id FROM Contact LIMIT 5]")
  }

  // Locate "[" and "]" character indexes
  firstOccurenceOfOpenSquare := strings.Index(fileContents, "[");
  firstOccurenceOfCloseSquare := strings.Index(fileContents, "]");

  if (firstOccurenceOfOpenSquare == -1 || firstOccurenceOfCloseSquare == -1) {
    log.Fatal("No query found in file, No query found in file, ensure the query is surrounded by square brackets. ie:\n[SELECT Id FROM Contact LIMIT 5]")
  }

  // Determine query to execute
  queryToExecute := fileContents[firstOccurenceOfOpenSquare+1:firstOccurenceOfCloseSquare];
  queryToExecute = strings.Replace(queryToExecute, "\n","",-1)


  return queryToExecute
}


func ReadBuildProperties() BuildPropertiesReadResult {
  buildPropertiesReadResult := readFile(BUILDPROPERTIES)
  buildProperties := strings.Split(buildPropertiesReadResult, "\n")

  return BuildPropertiesReadResult{
    Username: readFromBuildPropertiesFile(buildProperties, "sf.username = ", 0),
    Password: readFromBuildPropertiesFile(buildProperties, "sf.pass = ", 1),
    SecurityToken: readFromBuildPropertiesFile(buildProperties, "sf.securityToken = ", 2),
    ServerUrl: readFromBuildPropertiesFile(buildProperties, "sf.serverurl = ", 4),
    InstanceType: readFromBuildPropertiesFile(buildProperties, "sf.instancetype = ", 6),
  }
}

func readFromBuildPropertiesFile(buildProperties []string,
                                 searchPrefix string, 
                                 expectedIndex int) string {
  returnVal := ""
  length := len(buildProperties)
  indexExists := length >= (expectedIndex + 1)

  if !indexExists || !strings.HasPrefix(buildProperties[expectedIndex], searchPrefix) {
    log.Fatal(buildPropertiesReadKeyError(searchPrefix))
  } else {
    returnVal = strings.Replace(buildProperties[expectedIndex], searchPrefix, "", -1)
  }

  return returnVal
}