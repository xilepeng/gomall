package dal

import (
	"github.com/xilepeng/gomall/app/cart/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
