package httpserver
// Response
import (
	"fmt"
	"strings"
	"strconv"
)

// Response struct for an HTTP response.
type Response struct {
	StatusCode int
	Headers map[string]string
	Body []byte
}

var statusText = map[int]string{
	200: "OK",
	400: "Bad Request",
	404: "Not Found",
	500: "Internal Server Error",
}

func BuildResponse(resp *Response) []byte {
	var builder strings.Builder

	// Status
	reason := statusText[resp.StatusCode]
	if reason == "" {
		reason = "Unknown"
	}
	builder.WriteString(fmt.Sprintf("HTTP/1.1 %d %s\r\n", resp.StatusCode, reason))

	// Headers
	if resp.Headers == nil {
		resp.Headers = make(map[string]string)
	}
	resp.Headers["Content-Length"] = strconv.Itoa(len(resp.Body))

	for k, v := range resp.Headers {
		builder.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}

	// Body
	builder.WriteString("\r\n")
	builder.Write(resp.Body)

	return []byte(builder.String())
}
