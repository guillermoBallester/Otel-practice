# Otel-practice

This repository serves as a practical demonstration of OpenTelemetry (OTel) integration within a Go application. It showcases how to instrument a Go service to collect telemetry data, including traces and metrics, and export them to observability backends using the OpenTelemetry Collector.

## Project Structure

- **main.go**: Contains the primary Go application code, including OpenTelemetry instrumentation.
- **Dockerfile**: Defines the Docker image for the Go application.
- **docker-compose.yml**: Configures Docker services for the application, OpenTelemetry Collector, and observability backends.
- **otel-collector-config.yaml**: Configuration file for the OpenTelemetry Collector.
- **prometheus.yml**: Configuration file for Prometheus, used for metrics scraping.

## Prerequisites

Ensure you have the following installed:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/)

## Setup and Running the Application

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/guillermoBallester/Otel-practice.git
   cd Otel-practice
   ```

2. **Build the Go Application**:

   ```bash
   go build -o app main.go
   ```

3. **Run Using Docker Compose**:

   Start the application along with the OpenTelemetry Collector and observability backends:

   ```bash
   docker-compose up --build
   ```

   This command builds the Docker images and starts the services defined in `docker-compose.yml`.

## OpenTelemetry Instrumentation

The Go application is instrumented using the OpenTelemetry Go SDK to capture telemetry data. Key components include:

- **Tracer**: Captures trace data for operations within the application.
- **Meter**: Records metrics such as request counts and latencies.

Instrumentation is initialized in `main.go`, where the tracer and meter are configured, and spans are created to trace operations.

## OpenTelemetry Collector Configuration

The OpenTelemetry Collector is configured to receive telemetry data from the application and export it to various backends. The configuration (`otel-collector-config.yaml`) includes:

- **Receivers**: Accepts OTLP data over gRPC and HTTP.
- **Exporters**: Sends data to backends like Prometheus and Jaeger.
- **Processors**: Batches data before exporting.

## Observability Backends

The following observability backends are included:

- **Prometheus**: For metrics visualization.
- **Jaeger**: For trace visualization.

Access their UIs at:

- Prometheus: `http://localhost:9090`

## References

For more information on OpenTelemetry and its components:

- [OpenTelemetry Go SDK](https://github.com/open-telemetry/opentelemetry-go)
- [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector)
- [Prometheus](https://prometheus.io/)

---

This project provides a foundational example of integrating OpenTelemetry into a Go application, demonstrating the collection and export of telemetry data to observability tools.


