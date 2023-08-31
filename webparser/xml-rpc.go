package webparser

import (
	"apicheck/httplib"
	"encoding/xml"
	"strings"
)

type HTTPRequestResponse struct {
	Request  []byte
	Response []byte
}

type MethodCall struct {
	XMLName xml.Name `xml:"methodCall"`
	Method  string   `xml:"methodName"`
	Params  []Param  `xml:"params>param"`
}

type Param struct {
	Value Value `xml:"value"`
}

type Value struct {
	Data string `xml:",innerxml"`
}

func isXMLRPCCommunication(requestResponse httplib.Httpchat) bool {
	request := string(requestResponse.Httprequest)
	response := string(requestResponse.Httpresponse)

	// 判断请求和响应是否都符合 XML-RPC 格式
	return isXMLRPCMessage(request) && isXMLRPCMessage(response)
}

func isXMLRPCMessage(message string) bool {
	decoder := xml.NewDecoder(strings.NewReader(message))

	var call MethodCall
	err := decoder.Decode(&call)
	if err != nil {
		return false
	}

	return call.XMLName.Local == "methodCall" && call.Method != ""
}
