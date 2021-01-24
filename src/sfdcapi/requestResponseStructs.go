package sfdcapi

func loginXmlRequest(username string, password string) string {
  return `<?xml version="1.0" encoding="utf-8" ?>
    <env:Envelope xmlns:xsd="http://www.w3.org/2001/XMLSchema"
      xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
      xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
      <env:Body>
        <n1:login xmlns:n1="urn:partner.soap.sforce.com">
        <n1:username>` + username + `</n1:username>
        <n1:password>` + password + `</n1:password>
        </n1:login>
      </env:Body>
    </env:Envelope>`
}

// <?xml version="1.0" encoding="UTF-8"?>
// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns="urn:partner.soap.sforce.com" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
//   <soapenv:Body>
//     <loginResponse>
//       <result>
//         <metadataServerUrl></metadataServerUrl>
//         <passwordExpired></passwordExpired>
//         <sandbox></sandbox>
//         <serverUrl></serverUrl>
//         <sessionId></sessionId>
//         <userId></userId>
//         <userInfo>
//           <accessibilityMode></accessibilityMode>
//           <chatterExternal></chatterExternal>
//           <currencySymbol></currencySymbol>
//           <orgAttachmentFileSizeLimit></orgAttachmentFileSizeLimit>
//           <orgDefaultCurrencyIsoCode></orgDefaultCurrencyIsoCode>
//           <orgDefaultCurrencyLocale></orgDefaultCurrencyLocale>
//           <orgDisallowHtmlAttachments></orgDisallowHtmlAttachments>
//           <orgHasPersonAccounts></orgHasPersonAccounts>
//           <organizationId></organizationId>
//           <organizationMultiCurrency></organizationMultiCurrency>
//           <organizationName></organizationName>
//           <profileId></profileId>
//           <roleId></roleId>
//           <sessionSecondsValid></sessionSecondsValid>
//           <userDefaultCurrencyIsoCode xsi:nil="true"/>
//           <userEmail></userEmail>
//           <userFullName></userFullName>
//           <userId></userId>
//           <userLanguage></userLanguage>
//           <userLocale></userLocale>
//           <userName></userName>
//           <userTimeZone></userTimeZone>
//           <userType></userType>
//           <userUiSkin></userUiSkin>
//         </userInfo>
//       </result>
//     </loginResponse>
//   </soapenv:Body>
// </soapenv:Envelope>
type SoapLoginResponse struct {
  Body struct {
    LoginResponse struct {
      Result struct {
        MetadataServerUrl string `xml:"metadataServerUrl"`
        PasswordExpired bool `xml:"passwordExpired"`
        Sandbox bool `xml:"sandbox"`
        ServerUrl string `xml:"serverUrl"`
        SessionId string `xml:"sessionId"`
        UserId string `xml:"userId"`
        UserInfo struct {
          OrgDefaultCurrencyIsoCode string `xml:"orgDefaultCurrencyIsoCode"`
          OrgDefaultCurrencyLocale string `xml:"orgDefaultCurrencyLocale"`
          CurrencySymbol string `xml:"currencySymbol"`
          OrganizationId string `xml:"organizationId"`
          OrganizationName string `xml:"organizationName"`
          ProfileId string `xml:"profileId"`
          RoleId string `xml:"roleId"`
          UserEmail string `xml:"userEmail"`
          UserFullName string `xml:"userFullName"`
          UserId string `xml:"userId"`
          UserLanguage string `xml:"userLanguage"`
          UserLocale string `xml:"userLocale"`
          UserName string `xml:"userName"`
          UserTimeZone string `xml:"userTimeZone"`
          UserType string `xml:"userType"`
          UserUiSkin string `xml:"userUiSkin"`
          AccessibilityMode bool `xml:"accessibilityMode"`
          ChatterExternal bool `xml:"chatterExternal"`
          OrgDisallowHtmlAttachments bool `xml:"orgDisallowHtmlAttachments"`
          OrgHasPersonAccounts bool `xml:"orgHasPersonAccounts"`
          OrganizationMultiCurrency bool `xml:"organizationMultiCurrency"`
          OrgAttachmentFileSizeLimit float32 `xml:"orgAttachmentFileSizeLimit"`
          SessionSecondsValid float32 `xml:"sessionSecondsValid"`
        } `xml:"userInfo"`
      } `xml:"result"`
    } `xml:"loginResponse"`
  } `xml:"Body"`
}

