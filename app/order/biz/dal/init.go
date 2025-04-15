package dal

import (
	"github.com/xilepeng/gomall/app/order/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
