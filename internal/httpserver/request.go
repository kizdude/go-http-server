package httpserver

import (
	"fmt"
	"strings"
	"slices"
)
// Request Struct for parsed HTTP request.
type Request struct {
	Method string
	Path string
	Version string
	Headers map[string]string
	Body []byte
}

var methods = []string{
	"GET",
	"POST",
}

// ParseRequest takes int the data read over the stream and creates a request
// struct
func ParseRequest(data []byte) (*Request, error) {
	req := &Request{}

	requestStr := string(data)
	parts := strings.Split(requestStr, "\r\n\r\n")
	
	// Split Body from prefix
	if len(parts) != 2 {
		return nil, fmt.Errorf("malformed request")
	}

	// handle prefix things
	headerParts := strings.SplitN(parts[0], "\r\n", 2)

	requestLineParts := strings.Fields(headerParts[0])
	if len(requestLineParts) != 3 {
		return nil, fmt.Errorf("malformed request line")
	}

	if !slices.Contains(methods, requestLineParts[0]) {
		return nil, fmt.Errorf("invalid request method")
	}
	req.Method = requestLineParts[0]
	req.Path = requestLineParts[1]
	req.Version = requestLineParts[2]

	// headers
	headers := make(map[string]string)

	// If there are headers
	if len(headerParts) == 2 {
		headersParts := strings.Split(headerParts[1], "\r\n")
		// Populate headers map
		for i := range len(headersParts) {
			kv := strings.SplitN(headersParts[i], ":", 2)

			if len(kv) != 2 {
				return nil, fmt.Errorf("malformed header")
			}
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			headers[strings.ToLower(key)] = value
		}
	}

	req.Headers = headers

	req.Body = []byte(parts[1])

	return req, nil
}

func (req *Request) ToString() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Method: %s, Path: %s, Version: %s\n", req.Method, req.Path, req.Version))
	builder.WriteString("Headers:\n")
	for key, value := range req.Headers {
		builder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
	}
	builder.WriteString("Body: ")
	builder.Write(req.Body)

	return builder.String()
}
