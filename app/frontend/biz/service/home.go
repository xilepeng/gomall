package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xilepeng/gomall/app/frontend/infra/rpc"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Hot Sales",
		"items": products.Products,
	}, nil

	// var resp = make(map[string]any)

	// items := []map[string]any{
	// 	{"Name": "T-short-1", "Price": 100, "Picture": "/static/image/t-shirt-1.jpeg"},
	// 	{"Name": "T-short-2", "Price": 120, "Picture": "/static/image/t-shirt-2.jpeg"},
	// 	{"Name": "T-short-3", "Price": 130, "Picture": "/static/image/t-shirt-3.jpeg"},
	// 	{"Name": "T-short-4", "Price": 140, "Picture": "/static/image/t-shirt-4.jpeg"},
	// 	{"Name": "T-short-5", "Price": 150, "Picture": "/static/image/t-shirt-5.jpeg"},
	// 	{"Name": "T-short-6", "Price": 160, "Picture": "/static/image/t-shirt-6.jpeg"},
	// }
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items
	// return resp, nil

}

// var resp = make(map[string]any)

// items := []map[string]any{
// 	{"Name": "T-short-1", "Price": 100, "Picture": "/static/img/t-shirt-1.jpeg"},
// 	{"Name": "T-short-2", "Price": 120, "Picture": "/static/img/t-shirt-2.jpeg"},
// 	{"Name": "T-short-3", "Price": 130, "Picture": "/static/img/t-shirt-3.jpeg"},
// 	{"Name": "T-short-4", "Price": 140, "Picture": "/static/img/t-shirt-4.jpeg"},
// 	{"Name": "T-short-5", "Price": 150, "Picture": "/static/img/t-shirt-5.jpeg"},
// 	{"Name": "T-short-6", "Price": 160, "Picture": "/static/img/t-shirt-6.jpeg"},
// }
// resp["Title"] = "Hot Sales"
// resp["Items"] = items
// return resp, nil
