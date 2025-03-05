package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/xilepeng/gomall/app/frontend/conf"
	frontendUtils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		iniUserClient()
	})
}

func iniUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MuskHandleErr(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MuskHandleErr(err)
}
