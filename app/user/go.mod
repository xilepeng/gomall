module github.com/xilepeng/gomall/app/user

go 1.24.0

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/xilepeng/gomall/rpc_gen/kitex_gen => ../../rpc_gen
)

require gorm.io/gorm v1.25.12

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.14.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
