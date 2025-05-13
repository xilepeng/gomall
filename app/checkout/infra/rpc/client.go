package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xilepeng/gomall/app/checkout/conf"
	checkoututils "github.com/xilepeng/gomall/app/checkout/utils"
	"github.com/xilepeng/gomall/common/clientsuite"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	Once          sync.Once
	err           error
	serviceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
)

var commonOpts []client.Option

func InitClient() {
	Once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       RegistryAddr,
			// DestServiceAddr:    RegistryAddr,
			// TracerProvider:     mtl.TracerProvider,
		}),
	)

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: serviceName,
		RegistryAddr:       RegistryAddr,
		// TracerProvider:     mtl.TracerProvider,
	}))

	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: serviceName,
			RegistryAddr:       RegistryAddr,
			// TracerProvider:     mtl.TracerProvider,
		}),
	)
	opts = append(opts, commonOpts...)
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
	checkoututils.ShouldHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: serviceName,
		RegistryAddr:       RegistryAddr,
		// TracerProvider:     mtl.TracerProvider,
	}))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	opts = append(opts, commonOpts...)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoututils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r), client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName: serviceName,
		RegistryAddr:       RegistryAddr,
		// TracerProvider:     mtl.TracerProvider,
	}))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	opts = append(opts, commonOpts...)
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoututils.MustHandleError(err)
}
