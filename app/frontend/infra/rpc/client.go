package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xilepeng/gomall/app/frontend/conf"
	frontendutils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendutils.MustHandleError(err)
}

// func initProductClient() {
// 	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
// 	frontendutils.MustHandleError(err)
// 	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
// 	frontendutils.MustHandleError(err)
// }

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendutils.MustHandleError(err)
}
