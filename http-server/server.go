// 一个HTTP服务器
// 1. 接收客户端request，并将request中带的header写入response header
// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
// 3. Server端记录访问日志包括客户端ip, HTTP返回码, 输出到server端的标准输出
// 4. 当访问localhost/healthz时, 应返回200

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Server http-server定义
type Server struct {
	// 监听ip
	host string
	// 监听端口
	port int
	// 服务版本
	version string
	// 健康检测路径
	healthCheckPattern string
}

type ResponseWrapper interface {
	wrap(http.ResponseWriter, *http.Request)
}

type RequestHandler interface {
	handle(http.ResponseWriter, *http.Request)
}

// 初始化服务
func (s Server) init() {

	const V = "VERSION"
	version := os.Getenv(V)
	if version == "" {
		version = "1.0.0"
		err := os.Setenv(V, version)
		if err != nil {
			fmt.Println("server start fail：", err.Error())
		}
	}
	// 路由监听
	http.HandleFunc(s.healthCheckPattern, s.healthz)

	// 服务监听
	err := http.ListenAndServe(s.host+":"+strconv.Itoa(s.port), s)

	if err != nil {
		fmt.Println("server start fail：", err.Error())
	}

}

// 处理请求
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer s.responseWrapper(w, r)
}

// 响应处理
func (s Server) responseWrapper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("VERSION", s.version)
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
	}
}

// health check method
func (s Server) healthz(w http.ResponseWriter, r *http.Request) {
	writeString, err := io.WriteString(w, "200\n")
	if err != nil {
		fmt.Printf("health check error %d", writeString)
	}
}
