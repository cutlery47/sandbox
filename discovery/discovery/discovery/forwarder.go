package discovery

import (
	"discovery/common"
	"fmt"
	"net"
	"net/rpc"
)

type Forwarder interface {
	forward(*Connections, chan<- error)
	socket() *net.TCPAddr
}

type TCPMessageHandler struct {
	//TODO: figure out how to pass connections directly into handler functions
	conns    *Connections
	chin_err chan<- error
}

func (h TCPMessageHandler) HandleMessage(arg common.TCPMessage, reply *common.TCPMessage) error {
	*reply = common.TCPMessage{Data: "No response"}
	for _, conn := range *h.conns {
		err := conn.rpc.Call("TCPMessageHandler.HandleMessage", arg, reply)
		if err != nil {
			h.chin_err <- fmt.Errorf("couldn't send RPC request to %v", conn.addr.String())
		} else {
			break
		}
	}

	return nil
}

func (h TCPMessageHandler) HandleAnotherMessage(arg common.TCPMessage, reply *common.TCPMessage) error {
	*reply = common.TCPMessage{Data: "No response"}
	for _, conn := range *h.conns {
		err := conn.rpc.Call("TCPMessageHandler.HandleAnotherMessage", arg, reply)
		if err != nil {
			h.chin_err <- fmt.Errorf("couldn't send RPC request to %v", conn.addr.String())
		} else {
			break
		}
	}

	return nil
}

// implement Forwarder
type ForwardServer struct {
	sock    common.TCPSocket
	handler TCPMessageHandler
	l       net.Listener
	rpc     *rpc.Server
}

func (f ForwardServer) forward(conns *Connections, chin_err chan<- error) {
	for {
		conn, err := f.l.Accept()
		if err != nil {
			chin_err <- err
		}

		f.handler.chin_err = chin_err
		f.rpc.ServeConn(conn)
	}
}

func (f ForwardServer) socket() *net.TCPAddr {
	return &f.sock.SockAddr
}

func newForwarder(sock common.TCPSocket, conns *Connections) (ForwardServer, error) {
	l, err := net.Listen("tcp", sock.SockAddr.String())
	if err != nil {
		return ForwardServer{}, nil
	}

	forwarder := new(ForwardServer)
	forwarder.sock = sock
	forwarder.handler = TCPMessageHandler{}
	forwarder.handler.conns = conns

	forwarder.l = l
	forwarder.rpc = rpc.NewServer()
	forwarder.rpc.Register(forwarder.handler)

	return *forwarder, nil
}
