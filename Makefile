.PHONY: gen-demo-proto

gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC --server_name demo_proto --module github.com/xilepeng/gomall/demo/demo_proto -I ../../idl --idl ../../idl/echo.proto

gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --server_name demo_thrift --module github.com/xilepeng/gomall/demo/demo_thrift -I ../../idl --idl ../../idl/echo.thrift

