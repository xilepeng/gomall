.PHONY: gen-demo-proto

gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC -I ../../idl  --service demo_proto --module github.com/xilepeng/gomall/demo/demo_proto --idl ../../idl/echo.proto

gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --module github.com/xilepeng/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
