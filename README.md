# ⚡ Shoplite Backend API — Go (Golang) Microservice

[![Go CI Pipeline](https://github.com/Sais-gallego/shoplite-backend-api/actions/workflows/go-ci.yml/badge.svg)](https://github.com/Sais-gallego/shoplite-backend-api/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/Sais-gallego/shoplite-backend-api)](https://goreportcard.com/report/github.com/Sais-gallego/shoplite-backend-api)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A robust, production-grade e-commerce shopping cart and pricing engine microservice implemented in **Go (Golang)**. Demonstrates clean idiomatic Go patterns, pointer receiver semantics, concurrency-safe unit test coverage, and RESTful API endpoints.

Developed by **Sergio Andrés Ípuz Saavedra** (*Técnico en Programación de Software - SENA | Harvard CS50x*).

---

## 🎯 Architecture & Features

- **Idiomatic Go Domain Models**: Clean `Item`, `Cart`, and `Summary` struct definitions with JSON tags.
- **Smart Quantity Management**: `AddItem` detects existing SKUs and increments quantities in-place without duplicating records.
- **Slice Manipulation**: `RemoveItem` implements fast in-memory slice resectioning `append(items[:i], items[i+1:]...)` with custom error handling.
- **Financial Calculation Engine**: Computes nested subtotal, dynamic discount rates (e.g. promotional codes), and regional sales tax (VAT) with 2-decimal floating-point rounding.
- **100% Automated Test Suite**: Comprehensive unit testing using `testing.T` covering positive, boundary, and error cases.
- **Continuous Integration (CI)**: Automated GitHub Actions pipeline running `go test -race` on every push.

---

## 📂 Project Structure

```
shoplite-backend-api/
├── .github/workflows/
│   └── go-ci.yml             # GitHub Actions CI workflow
├── api/
│   └── handlers.go           # REST HTTP controller handlers
├── cart/
│   ├── cart.go               # Domain logic (AddItem, RemoveItem, CalculateTotal)
│   └── cart_test.go          # Unit tests suite
├── go.mod                    # Go module descriptor
├── main.go                   # HTTP server entrypoint
└── README.md                 # Engineering documentation
```

---

## 🧪 Running Unit Tests

To run the automated test suite with verbose execution logs:

```bash
go test -v ./...
```

Sample output:

```text
=== RUN   TestAddItem_NewItem
--- PASS: TestAddItem_NewItem (0.00s)
=== RUN   TestAddItem_ExistingQuantityIncrement
--- PASS: TestAddItem_ExistingQuantityIncrement (0.00s)
=== RUN   TestRemoveItem_Success
--- PASS: TestRemoveItem_Success (0.00s)
=== RUN   TestRemoveItem_NotFound
--- PASS: TestRemoveItem_NotFound (0.00s)
=== RUN   TestCalculateTotal_DiscountAndTax
--- PASS: TestCalculateTotal_DiscountAndTax (0.00s)
PASS
ok      github.com/Sais-gallego/shoplite-backend-api/cart   0.004s
```

---

## 🌐 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/v1/cart` | Returns the current shopping cart state |
| `POST` | `/api/v1/cart` | Adds an item or increments its quantity |
| `GET` | `/api/v1/cart/checkout?discount=0.10&tax=0.19` | Computes subtotal, discount, tax, and final total |

### Sample POST `/api/v1/cart` Request:
```json
{
  "id": "prod-101",
  "name": "Logitech MX Master 3S",
  "price": 99.99,
  "quantity": 1
}
```

---

## 👨‍💻 Author

**Sergio Andrés Ípuz Saavedra**
- **GitHub**: [@Sais-gallego](https://github.com/Sais-gallego)
- **LinkedIn**: [linkedin.com/in/sais-game-5a9207325](https://www.linkedin.com/in/sais-game-5a9207325/)
- **Email**: `g46113026@gmail.com`
- **Location**: Neiva, Huila, Colombia

---

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.
