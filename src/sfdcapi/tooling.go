package sfdcapi

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/xml"
	"strings"
)

type SfdcConnection struct {
	BaseUrl string
	SessionId string
  Authenticated bool
}

func AuthenticateToSFDC(username string,
                        password string,
                        securityToken string,
                        serverurl string) SfdcConnection {

  // Strip trailing / if any
  serverurl = serverurl + "/services/Soap/u/44.0"

  passwordWithToken := password + securityToken

  loginXmlRequest := loginXmlRequest(username, passwordWithToken)


  client := &http.Client{}
  req, err := http.NewRequest("POST", serverurl, bytes.NewBufferString(loginXmlRequest))
  if err != nil {
    // handle error
  }

  req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
  req.Header.Set("SOAPAction", "login")

  resp, err := client.Do(req)
  defer resp.Body.Close()

  auth := false;
  if resp.Status == "200 OK" {
    auth = true
    // fmt.Println("**** WAS 200 **")
  } else {
    auth = false
    // fmt.Println("response Headers:", resp.Header)
    // fmt.Println("**** FAILED **")
  }

  body, _ := ioutil.ReadAll(resp.Body)

  /**
   *  Extract Session Id
   *  and toolingUrl - if sucessful login..
   */

  var logResponse SoapLoginResponse
  xml.Unmarshal(body, &logResponse)
  // fmt.Println(logResponse)


  sessionId := logResponse.Body.LoginResponse.Result.SessionId
  
  toolingUrl := logResponse.Body.LoginResponse.Result.MetadataServerUrl
  // ensure this split worked...
  toolingUrl = strings.SplitAfter(toolingUrl, ".com")[0]


  // exampleForExecuteAnon := "/services/data/v46.0/tooling/executeAnonymous/?anonymousBody=System.debug('Test')%3B"
  // exampleForQuery := "/services/data/v46.0/tooling/query/?q=select+id+,name+from+user+limit+1"
  
  // toolingUrl = toolingUrl + exampleForQuery

  conn := SfdcConnection{
  	BaseUrl: toolingUrl,
  	SessionId: sessionId,
    Authenticated: auth,
  }

  return conn
}