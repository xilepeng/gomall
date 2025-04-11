package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	// "github.com/xilepeng/gomall/app/cart/utils"
	"github.com/cloudwego/hertz/pkg/common/utils"
	checkout "github.com/xilepeng/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/xilepeng/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/xilepeng/gomall/app/frontend/utils"
	rpccheckout "github.com/xilepeng/gomall/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/xilepeng/gomall/rpc_gen/kitex_gen/payment"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Address: &rpccheckout.Address{
			StreetAddress: req.Street,
			City:          req.City,
			State:         req.Province,
			Country:       req.Country,
			ZipCode:       req.Zipcode,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          req.CardNum,
			CreditCardCvv:             req.Cvv,
			CreditCardExpirationYear:  req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":    "waiting",
		"redirect": "/checkout/result",
	}, nil

}
