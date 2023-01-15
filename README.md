## clipboardshare-client

共享剪切板客户端程序 支持win、mac双系统

使用时需配置conf目录及conf.yml文件

### 相关仓库

[clipboardshare-web](https://github.com/sxzhi799/clipboardshare-web)


[clipboardshare-server](https://github.com/sxzhi799/clipboardshare-server)

### 编译命令

#### mac/linux
`go build -ldflags="-s -w" -o clipboard main.go`
#### windows
`go build -ldflags="-s -w" -o clipboard.exe main.go`

