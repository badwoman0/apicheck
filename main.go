package main

import (
	"apicheck/httplib"
	"fmt"
)

func main() {
	a := httplib.HttpParer("另外一次入侵.pcap")
	// for _, v := range a {
	// 	i, j := webparser.Root(v)
	// 	if i {
	// 		fmt.Printf("%v%v是api流量,类型为%v\n", string(v.Httprequest), string(v.Httpresponse), j)
	// 	}

	// }

	for _, v := range a {
		fmt.Printf("%v", string(v.Httprequest))
		fmt.Printf("%v", string(v.Httpresponse))

	}

}
