package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xilepeng/gomall/app/frontend/conf"
	frontendUtils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MuskHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MuskHandleError(err)
}

// func initProductClient() {
// 	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
// 	frontendUtils.MuskHandleError(err)
// 	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
// 	frontendUtils.MuskHandleError(err)
// }

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MuskHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MuskHandleError(err)
}


