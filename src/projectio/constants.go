package projectio

const SRCFOLDER = "src"
const EXECUTEANONFOLDER = "executeAnonymous"
const QUERYFOLDER = "query"
const BUILDPROPERTIES = "build.properties"
const BUILDXML = "build.xml"

func buildProperties(username string,
                     password string,
                     serverurl string,
                     instancetype string) string {
  return "sf.username = " + username +
    "\nsf.password = " + password +
    "\nsf.serverurl = " + serverurl +
    "\nsf.maxPoll = 20" +
    "\nsf.instancetype = " + instancetype
}

func openUrl(username string,
             password string,
             instanceUrl string) string {
  return "#!/bin/bash" +
    "\nopen '" + instanceUrl + "/login.jsp?pw=" + password + "&un=" + username + "'"
}

func login(username string,
             password string,
             instanceType string) string {
  return "#!/bin/bash" +
    "\nforce login -u=" + username + " -p='" + password + "' -i=" + instanceType
}

func executeAnon(iStr string) string {
  return "System.debug('hello world " + iStr + "');"
}

func buildXml() string {
  return `<project name="SFDC" default="test" basedir="." xmlns:sf="antlib:com.salesforce">
    <property file="build.properties"/>
    <property environment="env"/>

    <condition property="sf.username" value=""> <not> <isset property="sf.username"/> </not> </condition>
    <condition property="sf.password" value=""> <not> <isset property="sf.password"/> </not> </condition>
    <condition property="sf.sessionId" value=""> <not> <isset property="sf.sessionId"/> </not> </condition>

    <taskdef resource="com/salesforce/antlib.xml" uri="antlib:com.salesforce">
        <classpath>
            <pathelement location="../ant-salesforce.jar" />
        </classpath>
    </taskdef>

    <target name="pull">
      <sf:retrieve 
        username="${sf.username}"
        password="${sf.password}"
        sessionId="${sf.sessionId}"
        serverurl="${sf.serverurl}"
        maxPoll="${sf.maxPoll}"
        retrieveTarget="` + SRCFOLDER + `"
        unpackaged="` + SRCFOLDER + `/package.xml"/>
    </target>

</project>`
}

func packageXml() string {
  return `<?xml version="1.0" encoding="UTF-8"?>
<Package xmlns="http://soap.sforce.com/2006/04/metadata">
    <types>
        <members>*</members>
        <name>ApexClass</name>
    </types>
    <types>
        <members>*</members>
        <name>ApexComponent</name>
    </types>
    <types>
        <members>*</members>
        <name>ApexPage</name>
    </types>
    <types>
        <members>*</members>
        <name>ApexTrigger</name>
    </types>
    <types>
        <members>*</members>
        <name>StaticResource</name>
    </types>
    <version>40.0</version>
</Package>`
}

func buildPropertiesReadKeyError(findingKey string) string {
  exampleBuildProperties := buildProperties("username", "password", "serverurl", "instancetype")

  return `Unable to locate "` + findingKey + `" within the build properties file.

Expecting file named: ` + BUILDPROPERTIES + `

Having the format:
` + exampleBuildProperties
}