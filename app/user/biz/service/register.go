package service

import (
	"context"

	"github.com/xilepeng/gomall/app/frontend/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/user/model"
	user "github.com/xilepeng/gomall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/errgo.v2/fmt/errors"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.

	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password and password confirm not equal")
	}
	passwordhashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("password hash failed")
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwordhashed),
	}
	err = model.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
