package httplib

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// 解析报文，将其分为header和body

type Httpmessage struct {
	Header []byte
	Body   []byte
}

type Httpresult struct {
	Request  Httpmessage
	Response Httpmessage
}

func HttpParerMain(path string) []Httpresult {
	var res []Httpresult
	var temp Httpresult
	for _, v := range HttpParer(path) {
		temp.Request.Header, temp.Request.Body, _ = splitHTTPMessage(v.Httprequest)
		temp.Response.Header, temp.Response.Body, _ = splitHTTPMessage(v.Httpresponse)
		res = append(res, temp)
	}
	return res
}

type Httpchat struct {
	Httprequest  []byte
	Httpresponse []byte
}

// 解析报文，将其输出为http的request和response
func HttpParer(filepath string) []Httpchat {
	handle, err := pcap.OpenOffline(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var httpchats []Httpchat
	var temphttpchat Httpchat

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	var currentPayload []byte

	for packet := range packetSource.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer == nil {
			continue
		}
		payload := appLayer.Payload()

		if isHTTPStart(payload) {
			// 如果已经有未处理完的 HTTP 报文，先处理它
			if len(currentPayload) > 0 {
				if classifyHTTPMessage(currentPayload) {
					temphttpchat.Httpresponse = currentPayload
				} else {
					temphttpchat.Httprequest = currentPayload
				}
			}
			//todo 这块未处理完的http报文怎么放置呢
			currentPayload = payload
		} else if isHTTPEnd(payload) {
			currentPayload = append(currentPayload, payload...)
			if classifyHTTPMessage(currentPayload) {
				temphttpchat.Httpresponse = currentPayload
			} else {
				temphttpchat.Httprequest = currentPayload
			}
			if temphttpchat.Httprequest != nil && temphttpchat.Httpresponse != nil {
				httpchats = append(httpchats, temphttpchat)
				temphttpchat.Httprequest, temphttpchat.Httpresponse = nil, nil
			}
			currentPayload = nil
		} else if len(currentPayload) > 0 {
			currentPayload = append(currentPayload, payload...)
		}

	}
	return httpchats
}

func splitHTTPMessage(data []byte) (header []byte, body []byte, err error) {
	crlfIndex := bytes.Index(data, []byte("\r\n\r\n"))
	if crlfIndex == -1 {
		return nil, nil, fmt.Errorf("CRLF not found, invalid HTTP format")
	}

	header = data[:crlfIndex+4]
	body = data[crlfIndex+4:]

	return header, body, nil
}

func isHTTPStart(payload []byte) bool {
	return strings.HasPrefix(string(payload), "GET ") ||
		strings.HasPrefix(string(payload), "POST ") ||
		strings.HasPrefix(string(payload), "HTTP/1.0 ") ||
		strings.HasPrefix(string(payload), "HTTP/1.1 ") ||
		// 添加更多 HTTP 方法和状态判断
		false
}

func isHTTPEnd(payload []byte) bool {
	return strings.HasSuffix(string(payload), "\r\n\r\n")
}

func classifyHTTPMessage(data []byte) bool {
	if len(data) < 4 {
		return false
	}
	return data[0] == 'H' && data[1] == 'T' && data[2] == 'T' && data[3] == 'P'
	//返回true则为响应，返回false则为请求
	//这个处理逻辑有点简陋
}
