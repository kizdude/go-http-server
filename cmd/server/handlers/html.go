package handlers

import (
	"log"
	"os"
	"strconv"
	"github.com/kizdude/go-http-server/internal/httpserver"
)

func HtmlHandler(req *httpserver.Request) *httpserver.Response {
	body, err := os.ReadFile("./serve" + req.Path)
	
	if err != nil {
		log.Println("Error reading file:", err)
	}

	return &httpserver.Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/html",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}
}
