package handlers

import (
	"strconv"
	"github.com/kizdude/go-http-server/internal/httpserver"
)

func EchoHandler(req *httpserver.Request) *httpserver.Response {
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
