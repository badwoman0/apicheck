package webparser

import (
	"apicheck/httplib"
)

func Root(requestResponse httplib.Httpchat) (bool, string) {
	if isSOAPCommunication(requestResponse) {
		return true, "soap"
	} else if isXMLRPCCommunication(requestResponse) {
		return true, "xml rpc"

	} else if isRESTfulCommunication(requestResponse) {
		return true, "restful"

	} else if isGraphQLCommunication(requestResponse) {
		return true, "GraphQL"

	} else if isGRPCCommunication(requestResponse) {
		return true, "Grpc"

	} else {
		return false, ""
	}

}
