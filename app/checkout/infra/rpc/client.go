package rpc

import (
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
)
