package mtl

import (
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

var TracerProvider *tracesdk.TracerProvider

// func InitTracing(serviceName string) {
// 	exporter, err := otlptracegrpc.New(context.Background())
// 	if err != nil {
// 		panic(err)
// 	}
// 	server.RegisterShutdownHook(func() {
// 		exporter.Shutdown(context.Background()) //nolint:errcheck
// 	})
// 	processor := tracesdk.NewBatchSpanProcessor(exporter)
// 	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(serviceName)))
// 	if err != nil {
// 		res = resource.Default()
// 	}
// 	TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
// 	otel.SetTracerProvider(TracerProvider)
// }

func InitTracing(serviceName string) provider.OtelProvider {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithInsecure(),
		provider.WithEnableMetrics(false),
	)
	return p
}
