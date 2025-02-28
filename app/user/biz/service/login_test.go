package service

import (
	"context"
	"testing"

	"github.com/xilepeng/gomall/app/user/biz/dal/mysql"
	user "github.com/xilepeng/gomall/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
