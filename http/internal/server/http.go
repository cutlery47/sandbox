package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
)

type JSON map[string]string

type HTTPServer struct {
	tcp     TCPServer
	parser  HTTPRequestParser
	builder HTTPResponseBuilder
}

func (server *HTTPServer) Serve() error {
	listener, err := server.tcp.getListener()
	if err != nil {
		return err
	}

	for {
		conn, err := server.tcp.getConnection(listener)
		if err != nil {
			return err
		}

		fmt.Printf("Received a new connection: %v\n", conn.RemoteAddr())

		go server.handleHTTPConnection(conn)
	}
}

func (server *HTTPServer) handleHTTPConnection(conn net.Conn) {
	for {
		msg, err := server.tcp.readMsg(conn)
		if err != nil {
			return
		}

		request, err := server.parser.Parse(msg)
		if err != nil {
			return
		}

		response := server.handleHTTPRequest(request)

		buff := bytes.Buffer{}
		response.Write(&buff)
		err = server.tcp.writeMsg(buff.Bytes(), conn)
		if err != nil {
			return
		}
	}
}

func (server *HTTPServer) handleHTTPRequest(request http.Request) http.Response {
	body := request.Body
	return http.Response{Status: "200 OK", StatusCode: 200, Body: body}
}

type HTTPRequestParser struct{}

func (parser *HTTPRequestParser) Parse(msg []byte) (http.Request, error) {
	request, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(msg)))
	if err != nil {
		fmt.Println(err.Error())
		return http.Request{}, err
	}
	return *request, err
}

type HTTPResponseBuilder struct {
}

func (builder *HTTPResponseBuilder) Build(msg []byte) http.Response {
	return http.Response{}
}

func NewHTTPServer(host string, port int) (HTTPServer, error) {
	tcp, err := NewTCPServer(host, port)
	if err != nil {
		return HTTPServer{}, err
	}

	parser := HTTPRequestParser{}
	builder := HTTPResponseBuilder{}

	return HTTPServer{tcp: tcp, parser: parser, builder: builder}, nil
}
