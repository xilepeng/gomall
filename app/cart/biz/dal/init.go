package dal

import (
	"github.com/xilepeng/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
