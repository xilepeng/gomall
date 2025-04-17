package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xilepeng/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/order"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	
	return
}
