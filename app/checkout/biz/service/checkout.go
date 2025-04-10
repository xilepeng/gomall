package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/xilepeng/gomall/app/checkout/infra/rpc"

	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/xilepeng/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/payment"
	"github.com/xilepeng/gomall/rpc_gen/kitex_gen/product"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// 计算商品总价
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewBizStatusError(5004001, "cart is empty")
	}

	var total float32 // 商品总价
	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})
		if resultErr != nil {
			return nil, resultErr
		}
		if productResp == nil {
			continue
		}
		p := productResp.Product.Price
		cost := p * float32(cartItem.Quantity)
		total += cost
	}

	// 创建订单id
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()

	// 构建支付请求参数
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})

	if err != nil {
		klog.Error(err.Error())
	}

	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)

	if err != nil {
		return nil, err
	}

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return

}
