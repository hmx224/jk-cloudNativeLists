package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

/**
内容：编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把1都做完

1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回200
提交链接🔗：https://jinshuju.net/f/PlZ3xg
截止时间：10月7日晚23:59前
提示💡：
1、自行选择做作业的地址，只要提交的链接能让助教老师打开即可
2、自己所在的助教答疑群是几组，提交作业就选几组

*/

var (
	logger *zap.Logger
)

func init() {
	logger, _ = zap.NewProduction()
	logger.Warn("init ing")
}

func main() {
	http.HandleFunc("/request", request)

	http.HandleFunc("/getEnv", getEnv)

	http.HandleFunc("/getLog", getLog)

	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("端口监听失败：", err)
		return
	}
}

func request(writer http.ResponseWriter, request *http.Request) {
	logger.Info("进行request请求头获取")
	for k, v := range request.Header {
		//request 写入到 response 中
		for _, vv := range v {
			writer.Header().Add(k, vv)
		}

		//response 渲染到浏览器中
		_, err := io.WriteString(writer, fmt.Sprintf("%s : %s\n", k, v))
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(writer.Header())

}

func getEnv(writer http.ResponseWriter, request *http.Request) {
	logger.Info("进行环境变量获取")
	//获取系统所有环境变量
	for _, v := range os.Environ() {
		_, err := writer.Write([]byte(v + "\n"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	goversion := os.Getenv("GOVERSION")
	fmt.Println(goversion)

	writer.Header().Add("version", runtime.Version())
	_, err := writer.Write([]byte(fmt.Sprintf("version=%v\n", runtime.Version())))
	if err != nil {
		panic(err)
	}

	fmt.Println(writer.Header())

}

func getLog(writer http.ResponseWriter, request *http.Request) {
	logger.Info("进行日志监控")

	resp, err := http.Get("http://127.0.0.1:8888/request")
	if err != nil {
		fmt.Println("HTTP请求错误：", err)
		return
	}

	logger.Info("resp客户端IP：" + strings.Split(resp.Request.Host, ":")[0])
	logger.Info("respHTTP状态码：" + strconv.Itoa(resp.StatusCode))

	_, err = io.WriteString(writer, fmt.Sprintf("resp状态码：%v \n", resp.StatusCode))
	_, err = io.WriteString(writer, fmt.Sprintf("resp:\nhost:%v \n", resp.Request.Host))
	fmt.Println("request header info:")
	_, err = io.WriteString(writer, "request header info:\n")
	for k, v := range request.Header {
		_, err := io.WriteString(writer, fmt.Sprintf("%s : %s\n", k, v[0]))
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	logger.Info("进行健康检测")
	_, err := io.WriteString(writer, "200")
	if err != nil {
		panic(err)
	}
}
