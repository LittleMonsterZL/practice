package main

import (
	"net/http"
	"os"
)

func HandConn(w http.ResponseWriter, r *http.Request) {
	// 1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}

	// 2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	w.Header().Set("VERSION", os.Getenv("VERSION"))

	// 3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	// 这个题目我不会写 o(╥﹏╥)o

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	// 4. 当访问 localhost/healthz 时，应返回200
	w.Write([]byte("200"))
}

func main() {
	http.HandleFunc("/", HandConn)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe("localhost:80", nil)

}
