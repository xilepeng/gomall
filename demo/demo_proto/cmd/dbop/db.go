package main

import (
	"github.com/joho/godotenv"
	"github.com/xilepeng/gomall/demo/demo_proto/biz/dal"
)

// 测试数据库操作
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	// CRUD

	// create
	// mysql.DB.Create(&model.User{Email: "x@163.com", Password: "123456"})

	// update
	// mysql.DB.Model(&model.User{}).Where("email = ?", "x@163.com").Update("password", "0000")

	// read	021-*/+63
	// var row model.User
	// mysql.DB.Model(&model.User{}).Where("email = ?", "x@163.com").First(&row)
	// fmt.Printf("row: %+v\n", row)

	// delete
	// mysql.DB.Where("email = ?", "x@163.com").Delete(&model.User{})
	// 强制删除
	// mysql.DB.Unscoped().Where("email = ?", "x@163.com").Delete(&model.User{})
}
