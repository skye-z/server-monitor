# BetaX Server Monitor

A server monitoring tool

## 开始开发

```shell
# 初始化 Go 模块
go mod init server-monitor
```

## 编译打包
```shell
# 直接编译当前平台
go build -o server-monitor -ldflags '-s -w'

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server-monitor -ldflags '-s -w'
# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o server-monitor -ldflags '-s -w'
```