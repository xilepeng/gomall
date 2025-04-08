package dal

import (
	"github.com/xilepeng/gomall/app/payment/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
