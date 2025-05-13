package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/xilepeng/gomall/app/frontend/conf"
	frontendutils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/common/clientsuite"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = frontendutils.ServiceName
	MetricsPort    = conf.GetConf().Hertz.MetricsPort
	RegistryAddr   = conf.GetConf().Hertz.RegistryAddr
	err            error
)

func InitClient() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {

	UserClient, err = userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initCartClient() {

	CartClient, err = cartservice.NewClient("cart", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))
	frontendutils.MustHandleError(err)
}
