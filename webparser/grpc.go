package webparser

import (
	"apicheck/httplib"
	"strings"
)

//todo 这个逻辑需要重构

func isGRPCCommunication(requestResponse httplib.Httpchat) bool {
	request := requestResponse.Httprequest
	response := requestResponse.Httpresponse

	// 判断请求和响应是否都符合gRPC格式
	return isGRPCRequest(request) || isGRPCResponse(response)
}

func isGRPCRequest(request []byte) bool {
	// 检查是否为gRPC请求标识
	return strings.HasPrefix(string(request), "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n")
}

func isGRPCResponse(response []byte) bool {
	// 检查是否为gRPC响应标识
	return strings.HasPrefix(string(response), "HTTP/2.0 200 \r\ncontent-type: application/grpc\r\n")
}
