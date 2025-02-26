# cwgo server  --type RPC  --idl hello.proto  --server_name hellotest --module {{your_module_name}} -I .

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC--idl ../../idl/echo.proto --server_name demo_proto --module github.com/xilepeng/gomall/demo/demo_proto -I ../../idl 

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --idl ../../idl/echo.thrift --server_name demo_thrift --module github.com/xilepeng/gomall/demo/demo_thrift -I ../../idl 

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/about.proto --server_name frontend --module github.com/xilepeng/gomall/app/frontend -I ../../idl



