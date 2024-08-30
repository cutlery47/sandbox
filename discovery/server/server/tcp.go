package server

import (
	"discovery/common"
	"fmt"
	"net"
	"net/rpc"
)

const MsgSize = 512

type Server struct {
	msgHandler    TCPMessageHandler
	healthHandler TCPHealthHandler

	discoverySock common.TCPSocket
	serverSock    common.TCPSocket

	rpcServer *rpc.Server
}

func (server *Server) Serve() error {
	fmt.Println("Started serving")

	conn, err := net.DialTCP("tcp", &server.serverSock.SockAddr, &server.discoverySock.SockAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	server.rpcServer.ServeConn(conn)

	return nil

}

func New(host string, port string, discoveryHost string, discoveryPort string) (Server, error) {
	sock, err := common.NewTCPSocket(host, port)
	if err != nil {
		return Server{}, err
	}

	discoverySock, err := common.NewTCPSocket(discoveryHost, discoveryPort)
	if err != nil {
		return Server{}, err
	}

	msgHandler := new(TCPMessageHandler)
	healthHandler := new(TCPHealthHandler)

	rpc := rpc.NewServer()
	rpc.Register(msgHandler)
	rpc.Register(healthHandler)

	return Server{
		serverSock:    sock,
		msgHandler:    *msgHandler,
		healthHandler: *healthHandler,
		rpcServer:     rpc,
		discoverySock: discoverySock,
	}, nil
}
