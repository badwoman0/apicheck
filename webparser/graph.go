package webparser

import (
	"apicheck/httplib"
	"strings"
)

func isGraphQLCommunication(requestResponse httplib.Httpchat) bool {
	request := string(requestResponse.Httprequest)
	response := string(requestResponse.Httpresponse)

	// 判断请求和响应是否都符合GraphQL格式
	return isGraphQLRequest(request) || isGraphQLResponse(response)
}

func isGraphQLRequest(request string) bool {
	// 这里只是一个示例，可以根据实际情况进行修改
	return strings.Contains(request, "query") && strings.Contains(request, "{")
}

func isGraphQLResponse(response string) bool {
	// 这里只是一个示例，可以根据实际情况进行修改
	return strings.Contains(response, "{") && strings.Contains(response, "}")
}
