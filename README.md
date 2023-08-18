# Packer API

The **Packer API** is a Go-based HTTP server designed to handle requests related to packing calculations. It exposes endpoints for updating pack sizes and calculating packs based on certain criteria.

## Introduction

The Packer API is built to facilitate packing calculations, allowing users to update pack sizes and calculate packs based on the provided quantities. The API is implemented using Go's standard `net/http` package and follows a RESTful design for its endpoints.

## Installation

To use the Packer API, follow these steps:

1. **Prerequisites**: Ensure you have Go installed.

2. **Clone the Repository**: Clone this repository to your local machine.

3. **Configure**: Modify `config.Config` in the `config` package as needed. See [Configuration](#configuration) for details.

4. **Build and Run**: Execute `StartServer` in `main.go` to start the HTTP server. Use the `-config` flag to specify the configuration file path, e.g., `-config=config.json`.

## Configuration

Before running the Packer API, configure its settings by editing `config.Config` values:

```json
{
  "http_port": 7070,
  "grpc_port": 7071,
  "packet_sizes": [250, 500, 1000, 2000, 5000]
}
```

## Usage

Interact with the API using **curl** commands:

### Update Pack Sizes

```bash
curl -X POST -H "Content-Type: application/json" -d '{"sizes": [10, 20, 30]}' http://localhost:7070/update-pack-sizes
```

### Calculate Packs

```bash
curl -X POST -H "Content-Type: application/json" -d '{"quantity": 100}' http://localhost:7070/calculate-packs
```

## Conditional gRPC Support

The Packer API supports conditional compilation for gRPC:

```bash
go build -tags grpc
```

**TODO :** Grpc implementation has not been done yet.