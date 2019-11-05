package projectio

import (
  "strings"
  "log"
)

type BuildPropertiesReadResult struct {
  Username string
  Password string
  ServerUrl string
  InstanceType string
}

func ReadBuildProperties() BuildPropertiesReadResult {
  buildPropertiesReadResult := readFile(BUILDPROPERTIES)
  buildProperties := strings.Split(buildPropertiesReadResult, "\n")

  return BuildPropertiesReadResult{
    Username: readFromBuildPropertiesFile(buildProperties, "sf.username = ", 0),
    Password: readFromBuildPropertiesFile(buildProperties, "sf.password = ", 1),
    ServerUrl: readFromBuildPropertiesFile(buildProperties, "sf.serverurl = ", 2),
    InstanceType: readFromBuildPropertiesFile(buildProperties, "sf.instancetype = ", 4),
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