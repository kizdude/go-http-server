package httpserver

import (
    "net"
    "sync"
	"fmt"
	"log"
	"strconv"
)


// Server Struct to hold server state.
type Server struct {
    Addr   	string
    Router 	map[string]func(*Request) *Response
    wg     	sync.WaitGroup
	ln		net.Listener
}

// NewServer constructor creates a new server.
func NewServer(addr string) *Server {
	// Return reference to server struct
    return &Server{
		Addr: addr,
		Router: make(map[string]func(*Request) *Response),
	}
}

// Handle registers a path with a handler function.
func (s *Server) Handle(path string, handler func(*Request) *Response) {
	s.Router[path] = handler
}

// ListenAndServe starts the server and accepts connections.
func (s *Server) ListenAndServe() error {
	// Create server socket listener
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.ln = ln
    defer ln.Close()

	log.Printf("Server listening on port: %s", s.Addr)

	// Accept incoming connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		s.wg.Add(1)
		// Concurrently handle connection
		go s.handleConnection(conn)
	}
}

// Close procedure to shut down the server.
func (s *Server) Close() {
	s.ln.Close()
	s.wg.Wait()
}

// handleConnection processes a single client connection.
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading:", err)
			return
		}

		// Only use the data that was read
		data := buf[:n]

		req, err := ParseRequest(data)

		// error parsing request
		if err != nil {
			log.Println("400 Bad Request:", err, req)
			sendBadRequest(conn)
			return
		}

		connection, ok := req.Headers["Connection"]
		if ok {
			if connection == "close" {
				break
			}
		}

		handler, ok := s.Router[req.Path]
		var resp *Response
		if ok {
			resp = handler(req)
		} else {
			// path not found
			log.Println("404 Invalid path in request:", req.Path, req)
			body := []byte("Not Found")
			resp = &Response{
				StatusCode: 404,
				Headers: map[string]string{
					"Content-Type": "text/plain",
					"Content-Length": strconv.Itoa(len(body)),
				},
				Body: body,
			}
		}

		s.writeResponse(conn, resp)
	}
}

func sendBadRequest(conn net.Conn) {
	body := []byte("400 Bad Request")
	resp := &Response{
		StatusCode: 400,
		Headers: map[string]string{
			"Content-Type": "text/plain",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}
	
	conn.Write(BuildResponse(resp))
}

// writeResponse writes an HTTP response to the connection.
func (s *Server) writeResponse(conn net.Conn, resp *Response) {
	if _, err := conn.Write(BuildResponse(resp)); err != nil {
		log.Println("Error writing response", err)
	}
}
