package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func main() {
	ctx := context.Background()
	fmt.Println("Starting OpenTelemetry Exporter...")

	// Create gRPC metric exporter
	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint("http://otel-collector:4318"),
	)
	if err != nil {
		log.Fatalf("Failed to create OTLP metric exporter: %v", err)
	}

	fmt.Println("OTLP Exporter successfully created!")

	// Create a meter provider
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
		metric.WithResource(resource.Empty()), // Can be modified for specific metadata
	)
	otel.SetMeterProvider(meterProvider)

	// Create a meter
	meter := otel.Meter("basic-metrics")

	// Create a counter
	counter, err := meter.Int64Counter("test_counter")
	if err != nil {
		log.Fatalf("Failed to create counter: %v", err)
	}

	// Send test values every second
	for i := 0; i < 10; i++ {
		counter.Add(ctx, 1)
		fmt.Println("Incremented test_counter")
		time.Sleep(1 * time.Second)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("Metrics sent successfully!")
}
