module github.com/xilepeng/gomall/app/payment

go 1.23.8

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/xilepeng/gomall/rpc_gen/kitex_gen => ../../rpc_gen // 自定义模块远端如果找不到，去本地查找
)



require (
	
	github.com/cloudwego/fastpb v0.0.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
