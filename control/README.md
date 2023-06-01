# BetaX Server Monitor Control

## 特性
- 高效: 处理链路短、性能开销低
- 快捷: 单一产出物、一键即可用
- 易用: 可视化完善、全流程引导


## 编译打包
```shell
# 直接编译当前平台
go build -o ../out/monitor-control -ldflags '-s -w'

# 交叉编译linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../out/monitor-control -ldflags '-s -w'
# 交叉编译windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../out/monitor-control -ldflags '-s -w'
```