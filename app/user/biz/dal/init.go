package dal

import (
	"github.com/xilepeng/gomall/app/user/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
