package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) uint32 {
	return ctx.Value(SessionUserId).(uint32)
}
