# Product Management API

A RESTful API built with Go and Fiber for managing products and inventory. This project provides endpoints to create, read, update, delete products, and adjust inventory levels, with data persistence in memory and initial data loaded from a JSON file.

## Table of Contents
- [Product Management API](#product-management-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
    - [Optional Configuration with `.env`](#optional-configuration-with-env)
  - [Running the Application](#running-the-application)
    - [Locally](#locally)
    - [With Docker](#with-docker)
  - [API Endpoints](#api-endpoints)
    - [Example Usage](#example-usage)
  - [Docker Deployment](#docker-deployment)
  - [API Specification](#api-specification)
    - [Service Configurations in Choreo (optional)](#service-configurations-in-choreo-optional)
      - [Load initial data in Choreo ( optional )](#load-initial-data-in-choreo--optional-)

## Features
- **CRUD Operations**: Manage products with create, read, update, and delete operations.
- **Inventory Management**: Adjust product stock quantities.
- **In-Memory Storage**: Stores data in memory with initial data loaded from a JSON file.
- **Configurable**: Uses environment variables for configuration (e.g., port, hostname, initial data path).
- **Docker Support**: Containerized deployment with a multi-stage Dockerfile.

## Prerequisites
- **Go**: Version 1.23 or higher
- **Docker**: For containerized deployment
- **Git**: To clone the repository

## Installation
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd product-mgt-api
   ```

2. **Install Dependencies**:
   ```bash
   go mod download
   ```

3. **Prepare Initial Data**:
   Create a `configs/initial_data.json` file in the project root with the following structure:
   ```json
   [
       {
           "id": "12345",
           "name": "Wireless Mouse",
           "description": "A comfortable wireless mouse with ergonomic design.",
           "price": 25.99,
           "category": "Electronics",
           "stock_quantity": 150,
           "created_at": "2025-01-15T10:00:00Z",
           "updated_at": "2025-03-01T08:00:00Z"
       },
       {
           "id": "67890",
           "name": "Bluetooth Speaker",
           "description": "Portable Bluetooth speaker with high-quality sound.",
           "price": 49.99,
           "category": "Electronics",
           "stock_quantity": 85,
           "created_at": "2025-02-05T12:00:00Z",
           "updated_at": "2025-03-01T08:30:00Z"
       }
   ]
   ```

## Configuration
The application uses environment variables for configuration. You can set these variables directly or use a `.env` file (requires `godotenv` integration—see [Optional Configuration](#optional-configuration)).

| Variable         | Description                     | Default Value               |
| ---------------- | ------------------------------- | --------------------------- |
| `HOSTNAME`       | Hostname for the API server     | `localhost`                 |
| `PORT`           | Port for the API server         | `8080`                      |
| `ENV`            | Environment (e.g., development) | `""` (empty string)         |
| `INIT_DATA_PATH` | Path to initial data JSON file  | `configs/initial_data.json` |
| `BASE_PATH`      | API base path                   | `api/v1`                    |

### Optional Configuration with `.env`
To use a `.env` file:
1. Install `godotenv`:
   ```bash
   go get github.com/joho/godotenv
   ```
2. Add to `config/config.go` (in an `init()` function):
   ```go
   import "github.com/joho/godotenv"
   func init() {
       if err := godotenv.Load(); err != nil {
           log.Printf("Warning: No .env file found: %v", err)
       }
   }
   ```

## Running the Application
### Locally
```bash
go run main.go
```
- The API will start on `http://localhost:8080` (or the configured `HOSTNAME` and `PORT`).
- You’ll see: `API Running on port: 8080`.

### With Docker
1. **Build the Docker Image**:
   ```bash
   docker build -t product-api .
   ```

2. **Run the Container**:
   ```bash
   docker run -p 8080:8080 -e INIT_DATA_PATH=/app/configs/initial_data.json product-api
   ```
   - Maps port 8080 on the host to 8080 in the container.
   - Sets `INIT_DATA_PATH` to match the file location in the image.

3. **Verify**:
   Check logs:
   ```bash
   docker logs <container-id>
   ```
   Should show: `API Running on port: 8080`.

## API Endpoints

The default base path: `/api/v1`

| Method | Endpoint                | Description               |
| ------ | ----------------------- | ------------------------- |
| GET    | `/products`             | List all products         |
| POST   | `/products`             | Create a new product      |
| GET    | `/products/{productId}` | Get a product by ID       |
| PUT    | `/products/{productId}` | Update a product          |
| DELETE | `/products/{productId}` | Delete a product          |
| POST   | `/inventory/adjustment` | Adjust inventory quantity |

### Example Usage
**Get All Products**:
```bash
curl -X GET 'http://localhost:8080/api/v1/products'
```

**Create a Product**:
```bash
curl -X POST 'http://localhost:8080/api/v1/products' \
-H 'Content-Type: application/json' \
-d '{"name": "Noise Cancelling Headphones", "description": "Over-ear headphones", "price": 120.00, "category": "Electronics", "stock_quantity": 75}'
```

## Docker Deployment
The project uses a multi-stage Dockerfile for efficient builds:
- **Build Stage**: Compiles the Go binary with `golang:1.23-alpine`.
- **Runtime Stage**: Runs the binary in a minimal `alpine` image as a non-root user (UID 10014).

**Dockerfile Highlights**:
- Copies `configs/initial_data.json` to `/app/configs/` in the container.
- Sets the working directory to `/app`.
- Runs as a non-root user for security.

**Build and Run**:
```bash
docker build -t product-api .
docker run -p 8080:8080 -e INIT_DATA_PATH=/app/configs/initial_data.json product-api
```

**Debugging**:
- Check build output:
  ```bash
  docker build -t product-api --progress=plain --no-cache .
  ```
- View container logs:
  ```bash
  docker logs <container-id>
  ```

## API Specification
The API is documented in OpenAPI 3.0 format. See [openapi.yaml](docs/openapi.yaml) for details.

### Service Configurations in Choreo (optional)

Refer [config.go](internal/config/config.go) file for the available configurations.

For more information on how to configure a service in Choreo, refer [Manage Configurations and Secrets](https://wso2.com/choreo/docs/deploy/devops/configs-and-secrets/) documentation.

#### Load initial data in Choreo ( optional )

1. Set environment variable by navigating to Choreo Deploy page `INIT_DATA_PATH=configs/initial_data.json`
2. Mount the file contents of `configs/initial_data.json` in the path specified in step 1.

See [initial_data.json](configs/initial_data.json) for a sample file.