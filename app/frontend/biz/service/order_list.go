package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xilepeng/gomall/app/frontend/infra/rpc"
	"github.com/xilepeng/gomall/app/frontend/types"
	frontendUtils "github.com/xilepeng/gomall/app/frontend/utils"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/order"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {

	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var (
			total float32
			items []types.OrderItem
		)
		for _, v := range v.Items {
			total += v.Cost
			i := v.Item
			ProductResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: i.ProductId})
			if err != nil {
				return nil, err
			}
			if ProductResp == nil || ProductResp.Product == nil {
				continue
			}
			p := ProductResp.Product
			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        v.Cost,
				Qty:         i.Quantity,
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:    v.OrderId,
			CreatedDate: created.Format("2006-01-02 15:04:05"),
			Cost:       total,
			Items:       items,
		})

	}
	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
