package reduction

import (
	"bufio"
	"fmt"
	"strings"
)

// 基于静态资源请求的流量过滤
// 基于协议/状态码的流量过滤
// restful解析模块
// xml-rpc解析模块
//
//
//
//
//

func isStaticResource(requestPath string) bool {
	// 常见静态资源文件扩展名
	staticExtensions := []string{".css", ".js", ".png", ".jpg", ".jpeg", ".gif", ".svg"}

	// 判断请求路径是否以静态资源扩展名结尾
	for _, extension := range staticExtensions {
		if strings.HasSuffix(requestPath, extension) {
			return true
		}
	}
	return false
}

func CheckStaticResources(httpRequestString string) {

	// 将字符串读入到Scanner中
	scanner := bufio.NewScanner(strings.NewReader(httpRequestString))

	// 解析请求行
	scanner.Scan()
	requestLine := scanner.Text()
	parts := strings.Split(requestLine, " ")
	if len(parts) >= 2 {
		requestMethod := parts[0]
		requestPath := parts[1]

		// 判断是否为静态资源请求
		if requestMethod == "GET" && isStaticResource(requestPath) {
			fmt.Println("This is a static resource request")
		} else {
			fmt.Println("This is not a static resource request")
		}
	} else {
		fmt.Println("Invalid request")
	}
}
