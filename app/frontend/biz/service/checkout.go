package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/xilepeng/gomall/app/checkout/infra/rpc"
	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
	frontendutils "github.com/xilepeng/gomall/app/frontend/utils"
	rpccart "github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/xilepeng/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	var itmes []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		ProductResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		if ProductResp.Product == nil {
			continue
		}
		p := ProductResp.Product
		itmes = append(itmes, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}
	return utils.H{
		"title": "checkout",
		"itmes": itmes,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil

}
