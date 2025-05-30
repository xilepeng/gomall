package model

import (
	"context"

	"gorm.io/gorm"
)

// 收件人
type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       string
}

type Order struct {
	gorm.Model
	OrderId    string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId     uint32      `gorm:"type:int(11)"`
	Consignee  Consignee   `gorm:"embedded"` // 可嵌入一个结构体
	OrderItems []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
