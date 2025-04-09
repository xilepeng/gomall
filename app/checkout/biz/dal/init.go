package dal

import (
	"github.com/xilepeng/gomall/app/checkout/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
