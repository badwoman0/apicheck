package webparser

import (
	"apicheck/httplib"
	"bytes"
	"encoding/xml"
)

func isSOAPCommunication(requestResponse httplib.Httpchat) bool {
	request := string(requestResponse.Httprequest)
	response := string(requestResponse.Httpresponse)

	// 判断请求和响应是否都符合SOAP格式
	return isSOAPMessage(request) && isSOAPMessage(response)
}

func isSOAPMessage(message string) bool {
	decoder := xml.NewDecoder(bytes.NewReader([]byte(message)))

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch token.(type) {
		case xml.StartElement:
			// 如果遇到 StartElement，检查是否包含 SOAP 的命名空间
			elem := token.(xml.StartElement)
			if elem.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" {
				return true
			}
		}
	}

	return false
}
