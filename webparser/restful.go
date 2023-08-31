package webparser

import (
	"apicheck/httplib"
	"strings"
)

func isRESTfulCommunication(requestResponse httplib.Httpchat) bool {
	request := string(requestResponse.Httprequest)
	response := string(requestResponse.Httpresponse)

	// 判断请求和响应是否都符合RESTful格式
	return isRESTfulRequest(request) && isRESTfulResponse(response)
}

func isRESTfulRequest(request string) bool {
	// 这里只是一个示例，可以根据实际情况进行修改
	return strings.Contains(request, "GET") || strings.Contains(request, "POST") || strings.Contains(request, "PUT") || strings.Contains(request, "DELETE")
}

func isRESTfulResponse(response string) bool {
	// 这里只是一个示例，可以根据实际情况进行修改
	return strings.Contains(response, "200 OK") || strings.Contains(response, "201 Created") || strings.Contains(response, "204 No Content")
}
