package webparser

import (
	"apicheck/httplib"
	"fmt"
)

func Root(requestResponse httplib.Httpchat) {
	if isSOAPCommunication(requestResponse) {
		fmt.Println("soap格式")
	} else if isXMLRPCCommunication(requestResponse) {
		fmt.Println("xml rpc")

	} else if isRESTfulCommunication(requestResponse) {
		fmt.Println("restful")

	} else if isGraphQLCommunication(requestResponse) {
		fmt.Println("GraphQL")

	} else if isSOAPCommunication(requestResponse) {
		fmt.Println("soap")

	} else if isGRPCCommunication(requestResponse) {
		fmt.Println("Grpc")

	}
}
