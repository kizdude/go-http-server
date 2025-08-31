# A HTTP/1.1 Server Implementation in Golang

A minimal HTTP server written in Go using raw TCP connections, without the `net/http` package.  
Supports routing, GET/POST requests, and basic HTTP/1.1 functionality.

## Features

- Minimal HTTP server implementation
- Request parsing (`ParseRequest`)
- Response building (`BuildResponse`)
- Route handling via `Server.Handle(path, handler)`
- 400 Bad Request and 404 Not Found handling
- Supports multiple requests per TCP connection using goroutines

## Project Structure

```
myhttpserver/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go          # Entry point
└── internal/
    └── httpserver/          # Server implementation
        ├── request.go
        ├── response.go
        └── server.go

```

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/kizdude/go-http-server.git
cd go-http-server
````

2. Run the server:

```bash
go run ./cmd/server
```

3. Access in your browser or via `curl`:

```bash
curl http://localhost:8080/
curl -X POST http://localhost:8080/echo -d "Hello"
```

## License

This project is licensed under the MIT License.
