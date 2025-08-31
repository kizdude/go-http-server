package main

import (
	"log"
	"strconv"
	"github.com/kizdude/go-http-server/internal/httpserver"
)

func main() {
	server := httpserver.NewServer(":8080")

	server.Handle("/", rootHandler)
	server.Handle("/echo", echoHandler)

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Handlers
func rootHandler(req *httpserver.Request) *httpserver.Response {
	body := []byte("Hello, world!")
	return &httpserver.Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}
}

func echoHandler(req *httpserver.Request) *httpserver.Response {
	body := []byte(req.Body)
	return &httpserver.Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/plain",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}
}
