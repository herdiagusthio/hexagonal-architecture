# Hexagonal Architecture (Clean Architecture) Reference Implementation

This repository serves as a **Gold Standard** implementation of the Hexagonal Architecture pattern in Go. It is designed to demonstrate how to build a maintainable, testable, and decoupled system by separating the Core Domain from external Infrastructure.

## 🏗️ The Architecture

The project is divided into three distinct layers:

### 1. The Core (Business/Domain)
Located in `/business`. This is the "Heart" of the application.
- **Zero Dependencies**: The core does not know about GORM, HTTP, or any other external library.
- **Rich Domain Models**: Uses Value Objects (e.g., `Email`, `Username`) to enforce business invariants at the type level.
- **Interfaces**: Defines the `Repository` interface, specifying *what* the system needs, not *how* it's implemented.

### 2. The Adapters (Infrastructure)
Located in `/api` (Inbound) and `/repository` (Outbound).
- **API Adapter**: Handles HTTP requests and translates them into domain calls.
- **Repository Adapter**: Implements the domain interface using GORM. It translates database-specific errors into domain-specific errors (Error Translation).

### 3. The Application (App)
Located in `/app`. This is the "Glue" that wires the adapters to the core using Dependency Injection.

---

## 💎 Key Engineering Principles Applied

### 🚫 Prevention of Primitive Obsession
Instead of using `string` for everything, we use **Value Objects**. 
- `type Email string` $\rightarrow$ ensures an email is always valid before it ever reaches the service layer.

### 🛡️ Error Translation
The core never sees a `gorm.ErrRecordNotFound`. The repository adapter translates this into a `user.ErrUserNotFound`. This allows the core to remain agnostic of the persistence technology.

### 🧪 Total Testability
By using interfaces and dependency injection, the business layer is 100% testable without a database.

## 🚀 Getting Started

### Installation
```bash
go mod download
```

### Running Tests
```bash
go test ./... -v -cover
```
