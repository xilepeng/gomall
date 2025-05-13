package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/xilepeng/gomall/app/cart/conf"
	cartutils "github.com/xilepeng/gomall/app/cart/utils"
	"github.com/xilepeng/gomall/common/clientsuite"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	ServiceName   = conf.GetConf().Kitex.Service
	RegistryAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
			// DestServiceAddr:    RegisterAddr,
			// TracerProvider:     mtl.TracerProvider,
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
