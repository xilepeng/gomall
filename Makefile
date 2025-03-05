# cwgo server  --type RPC  --idl hello.proto  --server_name hellotest --module {{your_module_name}} -I .

export ROOTMOD=github.com/xilepeng/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC--idl ../../idl/echo.proto --server_name demo_proto --module ${ROOTMOD}/demo/demo_proto -I ../../idl 

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --idl ../../idl/echo.thrift --server_name demo_thrift --module ${ROOTMOD}/demo/demo_thrift -I ../../idl 

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/home.proto --server_name frontend --module ${ROOTMOD}/app/frontend -I ../../idl


.PHONY: gen-user
gen-user:
	@cd app/user && cwgo server --type RPC -I ../../idl --idl ../../idl/user.proto --server_name user --module ${ROOTMOD}/app/user  --pass "-use ${ROOTMOD}/rpc_gen/kitex_gen"
	@cd rpc_gen && cwgo client --type RPC --idl ../idl/user.proto --server_name user --module ${ROOTMOD}/rpc_gen -I ../idl
