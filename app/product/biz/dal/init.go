package dal

import (
	"github.com/xilepeng/gomall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
