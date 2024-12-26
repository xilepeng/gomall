package dal

import (
	"github.com/xilepeng/gomall/demo/demo_thrift/biz/dal/mysql"
	"github.com/xilepeng/gomall/demo/demo_thrift/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
