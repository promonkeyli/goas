# goas - Go + OAS (OpenAPI Specification)

Pure OpenAPI 3.2 generator for Go.

Parses Go code comments to generate OpenAPI 3.2 specification JSON.

## Features

- **Global Annotations**: Configure Info, Servers, Tags, Security in `main.go`.
- **API Annotations**: Define Router, Params, Request Body, Responses in handler functions.
- **Type Resolution**: Automatically converts Go structs to OpenAPI Schemas.
- **Generics Support**: Full support for Go 1.18+ generic types (e.g., `Response[User]`).
- **Standard Library Parsing**: Supports standard library types like `time.Time`, `os.File`.

## Status

**Current Version**: 0.1.0 (Development)

- [x] Global Annotation Parsing
- [x] API Operation Parsing
- [x] Schema Generation (Generics support included)
- [x] CLI Tool

## Installation

### Method 1: Global Installation (CLI)

Install the command-line tool globally:

```bash
go install github.com/promonkeyli/goas/cmd/goas@latest
```

### Method 2: Project Dependency (Library)

Add as a dependency to your project:

```bash
go get github.com/promonkeyli/goas
```

## Usage

### 1. CLI Usage

If installed globally, run directly:

```bash
goas -dir ./cmd,./internal -output ./api
```

Or run via go run in your project:

```bash
go run github.com/promonkeyli/goas/cmd/goas -dir ./cmd,./internal -output ./api
```

### 2. Library Usage (Function Call)

You can also integrate goas directly into your Go code (e.g., in a generator script):

```go
package main

import (
    "log"
    "github.com/promonkeyli/goas/pkg/goas"
)

func main() {
    config := goas.Config{
        Dirs:   []string{"./cmd", "./internal"},
        Output: "./api",
    }

    if err := goas.Run(config); err != nil {
        log.Fatalf("Failed to generate OpenAPI: %v", err)
    }
}
```

## Quick Start

### 1. Global Configuration (main.go)

```go
package main

// @OpenAPI 3.2.0
// @Title Pet Store API
// @Version 1.0.0
// @Description Supports user management and file uploads.
// @Server http://localhost:8080/v1 name=dev Development
// @SecurityScheme ApiKeyAuth apiKey header X-API-KEY
func main() {
    // ...
}
```

### 2. Implementation (handler.go)

```go
package handler

// Response Generic Response
type Response[T any] struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    Data    T      `json:"data"`
}

// User Model
type User struct {
    ID   int    `json:"id" desc:"User ID"`
    Name string `json:"name" desc:"Username"`
}

// GetUser Get User Info
// @Summary Get User Info
// @Tags user
// @Param id path int true "User ID"
// @Success 200 {object} Response[User] "Success"
// @Router /users/{id} [get]
func GetUser() {
    // ...
}
```

## Flags

- `-dir`: Comma-separated list of directories to scan (recursive).
- `-output`: Output directory for `openapi.json`.

## Documentation

- [GOAS Annotation Specification](docs/GOAS_COMMOENT.md): Detailed guide on using goas annotations.
- [OpenAPI 3.2.0 Field Reference](docs/OAS_FIELD.md): Quick reference for OpenAPI 3.2.0 objects and fields.
- [OpenAPI Specification 3.2.0](docs/OAS_STANDARD.md): Full OpenAPI 3.2.0 Specification.