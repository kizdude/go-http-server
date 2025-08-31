package main

import (
	"log"
	"github.com/kizdude/go-http-server/internal/httpserver"
	"github.com/kizdude/go-http-server/cmd/server/handlers"
)

func main() {
	server := httpserver.NewServer(":8080")

	server.Handle("/", handlers.RootHandler)
	server.Handle("/echo", handlers.EchoHandler)
	server.Handle("/index.html", handlers.HtmlHandler)

	log.Println("Server starting on port :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

