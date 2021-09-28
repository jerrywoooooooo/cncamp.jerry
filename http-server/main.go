// 一个HTTP服务器
// 1. 接收客户端request，并将request中带的header写入response header
// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
// 3. Server端记录访问日志包括客户端ip, HTTP返回码, 输出到server端的标准输出
// 4. 当访问localhost/healthz时, 应返回200

package main

import (
	"fmt"
	"github.com/felixge/httpsnoop"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {

	server := MyServer{
		version: "1.0.0",
		port:    80,
	}

	routes := make(map[string]func(http.ResponseWriter, *http.Request))

	routes["/index"] = server.defaultIndexHandler
	// 4. 当访问localhost/healthz时, 应返回200
	routes["/healthz"] = server.defaultHealthCheck
	routes["/custom"] = func(writer http.ResponseWriter, request *http.Request) {
		_, err := io.WriteString(writer, "this is custom path")
		if err != nil {
			log.Println("server response fail ", err.Error())
		}
	}

	server.routes = routes
	server.init()
}

// MyServer http-server定义
type MyServer struct {
	// 监听端口
	port int
	// 服务版本
	version string
	// 路由信息
	routes map[string]func(http.ResponseWriter, *http.Request)
}

const V = "VERSION"

// 初始化服务
func (s MyServer) init() {
	log.Printf("starting server on port %d\n", s.port)
	// 设置环境变量
	err := os.Setenv(V, s.version)
	if err != nil {
		log.Printf("set environment variable %s err\n", V)
	}

	mux := http.NewServeMux()

	if s.routes == nil || len(s.routes) == 0 {
		mux.HandleFunc("/", s.defaultIndexHandler)
	} else {
		// 路由监听
		for route, handler := range s.routes {
			mux.HandleFunc(route, handler)
		}
	}
	// 服务监听
	err = http.ListenAndServe(":"+strconv.Itoa(s.port), logRequestHandler(mux))
	if err != nil {
		log.Println("server start fail ", err.Error())
	}
}

// 处理请求
func (s MyServer) defaultIndexHandler(w http.ResponseWriter, r *http.Request) {
	s.setResponseHeader(w, r)
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			_, err := io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
			if err != nil {
				log.Printf("default handler error")
			}
		}
	}
}

// health check method
func (s MyServer) defaultHealthCheck(w http.ResponseWriter, r *http.Request) {
	s.setResponseHeader(w, r)
	_, err := io.WriteString(w, "200")
	if err != nil {
		log.Printf("default health check error")
	}
}

// 响应Header处理
func (s MyServer) setResponseHeader(w http.ResponseWriter, r *http.Request) {
	//1. 接收客户端request，并将request中带的header写入response header
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, strings.Join(v, ","))
		}
	}
	//w.Header().Set("VERSION", s.version)
	// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
	w.Header().Set(V, os.Getenv(V))
}

// 获取请求ip
func getRemoteAddr(r *http.Request) string {
	hdr := r.Header
	ip := hdr.Get("X-Forwarded-For")
	if ip != "" {
		return strings.Split(ip, ",")[0]
	}
	ip = hdr.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}
	return ""
}

const ignorePath = "/favicon.ico"

func logRequestHandler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.String()
		if path != ignorePath {
			metrics := httpsnoop.CaptureMetrics(h, w, r)
			// 3. Server端记录访问日志包括客户端ip, HTTP返回码, 输出到server端的标准输出
			log.Printf("[ip:%s] %s %s,code=%d,cost=%s,writen=%d", getRemoteAddr(r), r.Method, path, metrics.Code, metrics.Duration, metrics.Written)
		}
	}
	// 用 http.HandlerFunc 包装函数，这样就实现了 http.Handler 接口
	return http.HandlerFunc(fn)
}
