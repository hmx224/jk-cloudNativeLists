# jk-cloudNativeLists
针对极客时间-云原生课程练习作业清单

## 2021-10-07
第一次作业：增加http目录，编写一个 HTTP 服务器，增加http相关知识点练习

## 2021-10-16
windows下构建linux可运行的二进制包


```goland
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

go build -o bin/amd64/httpserver httpserver.go
