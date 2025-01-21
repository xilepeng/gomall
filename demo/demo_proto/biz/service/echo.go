package service

import (
	"context"
	"fmt"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/kerrors"
	pbapi "github.com/xilepeng/gomall/demo/demo_proto/kitex_gen/pbapi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *pbapi.Request) (resp *pbapi.Response, err error) {
	// Finish your business logic.
	if req.Message == "error" {
		return nil, kerrors.NewGRPCBizStatusError(1004001, "client param error")
	}

	clientName, ok := metainfo.GetPersistentValue(s.ctx, "CLIENT_NAME")
	fmt.Println(clientName, ok)
	return &pbapi.Response{Message: req.Message}, nil
}
