package dal

import (
	"github.com/xilepeng/gomall/app/frontend/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
