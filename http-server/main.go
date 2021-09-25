// 一个HTTP服务器
// 1. 接收客户端request，并将request中带的header写入response header
// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
// 3. Server端记录访问日志包括客户端ip, HTTP返回码, 输出到server端的标准输出
// 4. 当访问localhost/healthz时, 应返回200

package main

import (
	"flag"
	"io"
	"net/http"
	//_ "net/http/pprof"
)

func main() {
	flag.Set("v", "4")
	http.HandleFunc("/healthz", healthz)

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}
