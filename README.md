
## zsh 设置 GOPATH 环境变量
```shell
brew install go

sudo vim ~/.zshrc

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

source ~/.zshrc

go env -w GO111MODULE='on'
go env -w GOPROXY='https://goproxy.cn/,direct'
```


## 安装 cwgo 工具
```shell
 go install github.com/cloudwego/cwgo@latest
# thriftgo 安装：
go install github.com/cloudwego/thriftgo@latest
```

## 在 Zsh 中支持 临时支持 Zsh 补全

```shell
mkdir autocomplete # You can choose any location you like
cwgo completion zsh > ./autocomplete/zsh_autocomplete
source ./autocomplete/zsh_autocomplete


```


```shell
cd  demo/demo_thrift  

cwgo server --type RPC --module github.com/xilepeng/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

cd demo_proto 
cwgo server  --type RPC  -I ../../idl  --server_name demo_proto --module github.com/xilepeng/gomall/demo/demo_proto --idl ../../idl/echo.proto
```
