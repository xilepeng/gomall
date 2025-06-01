package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"
	"github.com/xilepeng/gomall/app/frontend/conf"
	frontendutils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/common/clientsuite"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product"
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
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)
	consulClient, err := consul.NewClient(consul.Options{})
	ProductClient, err = productcatalogservice.NewClient("product", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}), client.WithCircuitBreaker(cbs), client.WithFallback(
		fallback.NewFallbackPolicy(
			fallback.UnwrapHelper(
				func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
					if err == nil {
						return resp, nil
					}
					methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
					if methodName != "ListProducts" {
						return resp, err
					}
					return &product.ListProductsResp{
						Products: []*product.Product{
							{
								Price:       6.6,
								Id:          3,
								Picture:     "/static/image/t-shirt.jpeg",
								Name:        "T-Shirt",
								Description: "CloudWeGo T-Shirt",
							},
						},
					}, nil
				}),
		),
	),
		client.WithSuite(consulclient.NewSuite("product", ServiceName, consulClient)),
	)

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
