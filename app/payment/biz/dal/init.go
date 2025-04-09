package dal

import (
	"github.com/xilepeng/gomall/app/payment/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
