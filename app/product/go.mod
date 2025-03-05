module github.com/xilepeng/gomall/app/product

go 1.23.0

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/xilepeng/gomall/rpc_gen/kitex_gen/product => ../../rpc_gen/kitex_gen/product/productcatalogservice
)

require github.com/golang/protobuf v1.5.4 // indirect
