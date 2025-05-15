# cwgo server  --type RPC  --idl hello.proto  --server_name hellotest --module {{your_module_name}} -I .

export ROOT_MOD=github.com/xilepeng/gomall

.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server --type RPC--idl ../../idl/echo.proto --server_name demo_proto --module ${ROOT_MOD}/demo/demo_proto -I ../../idl 

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server --type RPC --idl ../../idl/echo.thrift --server_name demo_thrift --module ${ROOT_MOD}/demo/demo_thrift -I ../../idl 

.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server --type HTTP --idl ../../idl/frontend/order_page.proto --server_name frontend --module ${ROOT_MOD}/app/frontend -I ../../idl
	
.PHONY: gen-user
gen-user:
	@cd rpc_gen && cwgo client --type RPC  --server_name user --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/user.proto
	@cd app/user && cwgo server --type RPC -I ../../idl --idl ../../idl/user.proto --server_name user --module ${ROOT_MOD}/app/user  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen"
		
.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --server_name product --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/product.proto 
	@cd app/product && cwgo server --type RPC --server_name product --module ${ROOT_MOD}/app/product  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --server_name cart --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/cart.proto 
	@cd app/cart && cwgo server --type RPC --server_name cart --module ${ROOT_MOD}/app/cart  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/cart.proto

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --server_name payment --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/payment.proto 
	@cd app/payment && cwgo server --type RPC --server_name payment --module ${ROOT_MOD}/app/payment  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/payment.proto

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --server_name checkout --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/checkout.proto 
	@cd app/checkout && cwgo server --type RPC --server_name checkout --module ${ROOT_MOD}/app/checkout  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/checkout.proto

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --server_name order --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/order.proto 
	@cd app/order && cwgo server --type RPC --server_name order --module ${ROOT_MOD}/app/order  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/order.proto



.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --server_name email --module ${ROOT_MOD}/rpc_gen -I ../idl --idl ../idl/email.proto 
	@cd app/email && cwgo server --type RPC --server_name email --module ${ROOT_MOD}/app/email  --pass "-use ${ROOT_MOD}/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/email.proto


.PHONY: tidy-all
tidy-all: 
	cd app/cart && go mod tidy
	cd app/checkout && go mod tidy
	cd app/email && go mod tidy
	cd app/frontend && go mod tidy
	cd app/order && go mod tidy
	cd app/payment && go mod tidy
	cd app/product && go mod tidy
	cd app/user && go mod tidy
	cd common && go mod tidy

.PHONY: build-frontend
build-frontend:
	docker build -f ./deploy/Dockerfile.frontend -t frontend:${v} .
	# make build-frontend v=v1.1.1

.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .
	# make build-svc svc=product v=v1.1.1
	# docker run