module github.com/xilepeng/gomall/app/order

go 1.23.8

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

replace github.com/xilepeng/gomall/app/rpc_gen => ../../rpc_gen

require (
	go.uber.org/zap v1.27.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)

require (
	github.com/cloudwego/kitex v0.11.3 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	go.opentelemetry.io/otel v1.25.0 // indirect
	go.opentelemetry.io/otel/trace v1.25.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20241120035129-55da83caab1b
	go.uber.org/multierr v1.10.0 // indirect
)
