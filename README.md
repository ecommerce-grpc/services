

# Ecommerce gRPC Services

Ecommerce gRPC services using Hexagonal Architecture.

## Overview
This repository provides a set of microservices for ecommerce functionality, built using Go and leveraging gRPC for service communication. The architecture follows the Hexagonal (Ports & Adapters) pattern, emphasizing a clean separation between core business logic and external technologies.

## Hexagonal Architecture
The Hexagonal Architecture (also known as Ports and Adapters) ensures that the core business logic remains independent of frameworks, databases, and external APIs. This allows for easier testing, flexibility in replacing external dependencies, and greater maintainability.

## Structure

 - Core (Domain):

Contains the business logic and domain entities.
Unaware of frameworks, databases, or gRPC details.

- Ports:

Interfaces defined in the core to express requirements for external operations (e.g., repositories, services).
Examples: PaymentPort, OrderRepositoryPort.

- Adapters (Primary & Secondary):

Implementations of the ports.
Primary adapters handle incoming requests (e.g., gRPC handlers).
Secondary adapters handle outbound operations (e.g., database adapters, external API clients).

 - Application Layer:

Orchestrates use cases and coordinates flow between ports and adapters.

Example Directory Layout

```sh
/core
  /domain        # Business entities and value objects
  /ports         # Interfaces for input (driving) and output (driven) ports
  /usecase       # Application use cases (business rules)
/adapters
  /grpc          # gRPC server implementation (primary adapter)
  /db            # Database access implementation (secondary adapter)
  /external      # Other external integrations
```

### Getting Started
Clone the repository.

### Build and run the services with Go.
Interact via gRPC clients.

### Features
Modular, testable service design
gRPC APIs for fast and language-agnostic communication
Easily extensible with new adapters or integrations

### Contributing
Contributions are welcome! Please open issues or pull requests for enhancements, bug fixes, or questions.

License
MIT

