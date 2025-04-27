package dal

import (
	"github.com/xilepeng/gomall/app/email/biz/dal/mysql"
	"github.com/xilepeng/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
