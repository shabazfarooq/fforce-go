package main

/****
  ******* 1) implement options
  ******* 2) implement query command
             a) takes in param for query file location
             b) read content
             c) auth to sfdc
						 d) pass in contents
 - This should be able to run by passing in a file, and having it auto detect what should be happening? and passing args?

-- Add chmod capabilities for openUrl
-- Create ./login file as well
-- in init function, use the url from authCreds for creating openUrl file


  implement logger - replace all log.Fatal with an error method in projectio/logger
  finish implementing init function (missing the SFDCAPI stuff)
  change "src" var to be src and not src2 in projectio/constants

  change command interface to use inheritance

  // Command to implement. Open salesforce URL and login
*/

import (
	"fmt"
	"os"

	"../src/commands"
)

func main() {
	fmt.Println("Running...")
	fmt.Println("")
	fmt.Println("")

	// Args
	var argsWithoutProg = os.Args[1:]
	// fmt.Println(argsWithoutProg)

	// add message for missing command..

	var commandRequested = argsWithoutProg[0]
	fmt.Println("commandRequested:" + commandRequested)

	// for i := 0; i < len(argsWithoutProg); i++ {
	// 	fmt.Println(argsWithoutProg[i])
	// }
	// first one should be command
	// if

	// Execute command
	var command commands.Command

	command = &commands.Init{}
	command.New("optionObject")

	// command = &commands.Query{}
	// command.New("optionObject")

	// // <?xml version="1.0" encoding="UTF-8"?>
	// // <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:partner.soap.sforce.com" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	// //   <soapenv:Body>
	// //     <loginResponse>
	// //       <result>
	// //         <metadataServerUrl></metadataServerUrl>
	// //         <passwordExpired></passwordExpired>
	// //         <sandbox></sandbox>
	// //         <serverUrl></serverUrl>
	// //         <sessionId></sessionId>
	// //         <userId></userId>
	// //         <userInfo>
	// //           <accessibilityMode></accessibilityMode>
	// //           <chatterExternal></chatterExternal>
	// //           <currencySymbol></currencySymbol>
	// //           <orgAttachmentFileSizeLimit></orgAttachmentFileSizeLimit>
	// //           <orgDefaultCurrencyIsoCode></orgDefaultCurrencyIsoCode>
	// //           <orgDefaultCurrencyLocale></orgDefaultCurrencyLocale>
	// //           <orgDisallowHtmlAttachments></orgDisallowHtmlAttachments>
	// //           <orgHasPersonAccounts></orgHasPersonAccounts>
	// //           <organizationId></organizationId>
	// //           <organizationMultiCurrency></organizationMultiCurrency>
	// //           <organizationName></organizationName>
	// //           <profileId></profileId>
	// //           <roleId></roleId>
	// //           <sessionSecondsValid></sessionSecondsValid>
	// //           <userDefaultCurrencyIsoCode xsi:nil="true"/>
	// //           <userEmail></userEmail>
	// //           <userFullName></userFullName>
	// //           <userId></userId>
	// //           <userLanguage></userLanguage>
	// //           <userLocale></userLocale>
	// //           <userName></userName>
	// //           <userTimeZone></userTimeZone>
	// //           <userType></userType>
	// //           <userUiSkin></userUiSkin>
	// //         </userInfo>
	// //       </result>
	// //     </loginResponse>
	// //   </soapenv:Body>
	// // </soapenv:Envelope>
	// type SoapLoginResponse struct {
	//   Body struct {
	//     LoginResponse struct {
	//       Result struct {
	//         MetadataServerUrl string `xml:"metadataServerUrl"`
	//         PasswordExpired bool `xml:"passwordExpired"`
	//         Sandbox bool `xml:"sandbox"`
	//         ServerUrl string `xml:"serverUrl"`
	//         SessionId string `xml:"sessionId"`
	//         UserId string `xml:"userId"`
	//         UserInfo struct {
	//           OrgDefaultCurrencyIsoCode string `xml:"orgDefaultCurrencyIsoCode"`
	//           OrgDefaultCurrencyLocale string `xml:"orgDefaultCurrencyLocale"`
	//           CurrencySymbol string `xml:"currencySymbol"`
	//           OrganizationId string `xml:"organizationId"`
	//           OrganizationName string `xml:"organizationName"`
	//           ProfileId string `xml:"profileId"`
	//           RoleId string `xml:"roleId"`
	//           UserEmail string `xml:"userEmail"`
	//           UserFullName string `xml:"userFullName"`
	//           UserId string `xml:"userId"`
	//           UserLanguage string `xml:"userLanguage"`
	//           UserLocale string `xml:"userLocale"`
	//           UserName string `xml:"userName"`
	//           UserTimeZone string `xml:"userTimeZone"`
	//           UserType string `xml:"userType"`
	//           UserUiSkin string `xml:"userUiSkin"`
	//           AccessibilityMode bool `xml:"accessibilityMode"`
	//           ChatterExternal bool `xml:"chatterExternal"`
	//           OrgDisallowHtmlAttachments bool `xml:"orgDisallowHtmlAttachments"`
	//           OrgHasPersonAccounts bool `xml:"orgHasPersonAccounts"`
	//           OrganizationMultiCurrency bool `xml:"organizationMultiCurrency"`
	//           OrgAttachmentFileSizeLimit float32 `xml:"orgAttachmentFileSizeLimit"`
	//           SessionSecondsValid float32 `xml:"sessionSecondsValid"`
	//         } `xml:"userInfo"`
	//       } `xml:"result"`
	//     } `xml:"loginResponse"`
	//   } `xml:"Body"`
	// }

	//  /*
	//  connect to tooling
	//  */
	// url := "https://test.salesforce.com/services/Soap/u/44.0"

	// loginXmlRequest := `<?xml version="1.0" encoding="utf-8" ?>
	//   <env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	//     xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	//     xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	//     <env:Body>
	//       <n1:login xmlns:n1="urn:partner.soap.sforce.com">
	//       <n1:username></n1:username>
	//       <n1:password></n1:password>
	//       </n1:login>
	//     </env:Body>
	//   </env:Envelope>`

	// client := &http.Client{}
	// req, err := http.NewRequest("POST", url, bytes.NewBufferString(loginXmlRequest))
	// if err != nil {
	//   // handle error
	// }

	// req.Header.Set("Content-Type", "text/xml; charset=UTF-8")
	// req.Header.Set("SOAPAction", "login")

	// resp, err := client.Do(req)
	// defer resp.Body.Close()

	// if resp.Status == "200 OK" {
	//   fmt.Println("**** WAS 200 **")
	// } else {
	//   fmt.Println("response Headers:", resp.Header)
	//   fmt.Println("**** FAILED **")
	// }

	// body, _ := ioutil.ReadAll(resp.Body)

	// // /**
	// //  *  Extract Session Id
	// //  *  and toolingUrl - if sucessful login..
	// //  */

	// var logResponse SoapLoginResponse
	// xml.Unmarshal(body, &logResponse)
	// // fmt.Println(logResponse)

	// sessionId := logResponse.Body.LoginResponse.Result.SessionId

	// toolingUrl := logResponse.Body.LoginResponse.Result.MetadataServerUrl
	// // ensure this split worked...
	// toolingUrl = strings.SplitAfter(toolingUrl, ".com")[0]

	// // exampleForExecuteAnon := "/services/data/v46.0/tooling/executeAnonymous/?anonymousBody=System.debug('Test')%3B"
	// exampleForQuery := "/services/data/v46.0/tooling/query/?q=select+id+,name+from+user+limit+1"

	// toolingUrl = toolingUrl + exampleForQuery

	// fmt.Println(sessionId)
	// fmt.Println(toolingUrl)
	// fmt.Println("")
	// fmt.Println("**** Completed logging in... *****")
	// fmt.Println("")

	// // /**
	// //  * Now, use JSON to do whatever tooling requests
	// //  */

	// client2 := &http.Client{}
	// req2, err2 := http.NewRequest("GET", toolingUrl, bytes.NewBufferString("{}"))
	// if err2 != nil {
	//   // handle error
	// }

	// req2.Header.Set("Content-Type", "application/json")
	// req2.Header.Set("Authorization", "Bearer " + sessionId)

	// resp2, err2 := client2.Do(req2)
	// defer resp2.Body.Close()

	// // fmt.Println("response Status:", resp2.Status)
	// // fmt.Println("response Headers:", resp2.Header)
	// body2, _ := ioutil.ReadAll(resp2.Body)
	// fmt.Println("response Body:", string(body2))

	// login using soap
	// extract the login url
	// use regular callout with Authorization: Bearer asdlkfjaskdlfj and Content-Type: application/json
	// /services/data/v20.0/query/?q
	// https://na85.salesforce.com/services/data/v44.0/tooling/query/?q=Select+name+from+account

}
