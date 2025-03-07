package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	frontendUtils "github.com/xilepeng/gomall/app/frontend/utils"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WarpResponse(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	// content["user_id"] = 222

	// session := sessions.Default(c)
	// userId := session.Get("user_id")
	// content["user_id"] = userId
	content["user_id"] = ctx.Value(frontendUtils.SessionUserId)
	return content

}
