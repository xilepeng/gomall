package service

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/xilepeng/gomall/app/user/biz/dal/mysql"
	user "github.com/xilepeng/gomall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "x@163.com",
		Password:        "0000",
		PasswordConfirm: "0000",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
