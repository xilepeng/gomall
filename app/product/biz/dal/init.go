package dal

import (
	"github.com/xilepeng/gomall/app/product/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
