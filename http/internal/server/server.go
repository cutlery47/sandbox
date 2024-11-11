package server

import (
	"fmt"
	"io"
	"net"
)

const MsgSize = 512

type TCPServer struct {
	sock    TCPServerSocket
	handler TCPMessageHandler
}

type TCPMessage []byte

type TCPMessageHandler struct{}

func (server *TCPServer) Serve() error {

	listener, err := server.getListener()
	if err != nil {
		return err
	}

	for {
		new_conn, err := server.getConnection(listener)
		fmt.Printf("Received a new connection: %v\n", new_conn.RemoteAddr())

		if err != nil {
			return err
		}

		go server.handleTCPConnection(new_conn)
	}
}

func (server *TCPServer) handleTCPConnection(conn net.Conn) {
	for {
		msg, err := server.readMsg(conn)
		if err != nil {
			return
		}

		response := server.handler.handle(msg)

		err = server.writeMsg(response, conn)
		if err != nil {
			return
		}
	}
}

func (server *TCPServer) readMsg(conn net.Conn) (TCPMessage, error) {
	msg := make([]byte, MsgSize)

	// Reading bytes from socket
	read, err := conn.Read(msg)
	fmt.Printf("Bytes read: %v\nContent: %c\nRead error: %v\n", read, msg, err)

	if err == io.EOF {
		conn.Close()
		fmt.Printf("Connection %v closed\n", conn.RemoteAddr())
		return nil, err
	}

	return TCPMessage(msg), err
}

func (server *TCPServer) writeMsg(msg TCPMessage, conn net.Conn) error {
	wrote, err := conn.Write(msg)
	fmt.Printf("Wrote %v bytes. Write error: %v\n", wrote, err)

	if err != nil {
		fmt.Printf("Error when writing to %v. Error: %v\n", conn.RemoteAddr(), err.Error())
	}

	return err
}

func (server *TCPServer) getListener() (net.Listener, error) {
	listener, err := net.Listen(server.sock.SockAddr.Network(), server.sock.SockAddr.String())
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func (server *TCPServer) getConnection(listener net.Listener) (net.Conn, error) {
	conn, err := listener.Accept()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (handler *TCPMessageHandler) handle(msg TCPMessage) TCPMessage {
	response := "Echo: " + string(msg)
	return TCPMessage([]byte(response))
}

func NewTCPServer(host string, port int) (TCPServer, error) {
	sock, err := newTCPServerSocket(host, port)
	if err != nil {
		return TCPServer{}, err
	}

	return TCPServer{sock: sock}, nil
}
