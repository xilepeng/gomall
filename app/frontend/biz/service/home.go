package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)

	items := []map[string]any{
		{"Name": "T-short-1", "Price": 100, "Picture": "/static/img/t-shirt-1.jpeg"},
		{"Name": "T-short-2", "Price": 120, "Picture": "/static/img/t-shirt-2.jpeg"},
		{"Name": "T-short-3", "Price": 130, "Picture": "/static/img/t-shirt-3.jpeg"},
		{"Name": "T-short-4", "Price": 140, "Picture": "/static/img/t-shirt-4.jpeg"},
		{"Name": "T-short-5", "Price": 150, "Picture": "/static/img/t-shirt-5.jpeg"},
		{"Name": "T-short-6", "Price": 160, "Picture": "/static/img/t-shirt-6.jpeg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items

	return resp, nil
}
