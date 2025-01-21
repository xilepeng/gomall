package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	api "github.com/xilepeng/gomall/demo/demo_thrift/kitex_gen/api"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *api.Request) (resp *api.Response, err error) {
	// Finish your business logic.
	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Println(info.From().ServiceName())

	// 忘记写导致返回值为空
	return &api.Response{Message: req.Message}, nil
}
