# Sezzle Calculator

Full-stack calculator built with React + TypeScript and Go.

The application supports:

- Addition
- Subtraction
- Multiplication
- Division

All calculations are performed by the Go backend through an HTTP API.

## Tech Stack

- **Frontend:** React, TypeScript, Vite
- **Backend:** Go
- **Testing:** Vitest, React Testing Library, Go testing package

## Running the Application

### Backend

```bash
cd backend
go run ./cmd/server
```

The API runs on:
`http://localhost:8080`

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Vite proxies `/api` requests to the backend during development.

## API

`POST /api/calculate`

Request:

```json
{
  "operation": "add",
  "operand1": 10,
  "operand2": 5
}
```

Supported operations:

- add
- subtract
- multiply
- divide

Successful response:

```json
{
  "result": 15
}
```

Example error response:

```json
{
  "error": "division by zero is not allowed"
}
```

## Testing

### Backend

```bash
cd backend
go test ./...
```

Coverage:

```bash
go test -coverprofile="./coverage.out" ./...
go tool cover -func="./coverage.out"
```

### Frontend

```bash
cd frontend
npm run test:run
```

Coverage:

```bash
npm run test:coverage
```

Additional checks:

```bash
npm run build
npm run lint
```

## Design Notes

- HTTP handling and calculation logic are separated in the backend.
- Request validation distinguishes missing operands from valid zero values.
- Frontend API communication is isolated from the calculator component.
- Arithmetic is performed only by the backend.
- Optional operations were intentionally excluded to prioritize correctness, clarity, and testing.

## AI Usage

AI tools were used during development for architecture discussion, code review, testing guidance, and debugging.
See [`AI_USAGE.md`](./AI_USAGE.md) for details.
