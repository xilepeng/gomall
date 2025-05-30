package checkout

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xilepeng/gomall/app/frontend/biz/service"
	"github.com/xilepeng/gomall/app/frontend/biz/utils"
	checkout "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/checkout"
	common "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/common"
)

// Checkout .
// @router /checkout [GET]
func Checkout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, hertzUtils.H{"watning": err}))
		return
	}

	resp, err := service.NewCheckoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "checkout", utils.WarpResponse(ctx, c, resp))
}

// CheckoutWaiting .
// @router /checkout/waiting [POST]
func CheckoutWaiting(ctx context.Context, c *app.RequestContext) {
	var err error
	var req checkout.CheckoutReq
	err = c.BindAndValidate(&req)
	if err != nil {
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	resp, err := service.NewCheckoutWaitingService(ctx, c).Run(&req)
	if err != nil {
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"error": err}))
		return
	}
	c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, resp))
}

// CheckoutResult .
// @router /checkout/result [GET]
func CheckoutResult(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}

	resp, err := service.NewCheckoutResultService(ctx, c).Run(&req)
	if err != nil {
		// utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		c.HTML(consts.StatusOK, "waiting", utils.WarpResponse(ctx, c, hertzUtils.H{"warning": err}))
		return
	}
	c.HTML(consts.StatusOK, "result", utils.WarpResponse(ctx, c, resp))
}
