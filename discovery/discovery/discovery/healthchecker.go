package discovery

import (
	"discovery/common"
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Healthchecker interface {
	register(*Connections, chan<- error)
	healthcheck(*Connections, chan<- error)
	socket() *net.TCPAddr
}

// implements Healthchecker
type HealthServer struct {
	sock common.TCPSocket
	l    net.Listener
}

func (h HealthServer) register(conns *Connections, chin_err chan<- error) {
	for {
		// listening for incoming connections from servers
		conn, err := h.l.Accept()
		if err != nil {
			conn.Close()
			chin_err <- err
			return
		}

		// converting recieved address to TCPAddr
		tcp_addr, err := net.ResolveTCPAddr("tcp", conn.RemoteAddr().String())
		if err != nil {
			conn.Close()
			chin_err <- err
			return
		}

		// creating and saving a new connection object
		*conns = append(*conns,
			TCPServer{
				conn: conn,
				addr: *tcp_addr,
				rpc:  rpc.NewClient(conn),
			})
	}
}

func (h HealthServer) healthcheck(conns *Connections, chin_err chan<- error) {
	var mu sync.Mutex
	for {
		mu.Lock()
		for i, conn := range *conns {
			reply := new(common.TCPMessage)
			// check if server is responsive
			err := conn.rpc.Call("TCPHealthHandler.HandleHealthcheck", common.TCPMessage{Data: "u ok?"}, &reply)
			if reply.Data != "ok" || err != nil {
				if err == nil {
					err = fmt.Errorf("got wrong response from %v", conn.addr.String())
				}
				h.handleHealthcheckError(conn.conn, conns, i, err, chin_err)
			}
		}

		mu.Unlock()
		// Healthcheck summary
		conns.printAll()
		time.Sleep(time.Millisecond * 500)
	}
}

func (h HealthServer) socket() *net.TCPAddr {
	return &h.sock.SockAddr
}

func (h HealthServer) handleHealthcheckError(conn net.Conn, conns *Connections, i int, err error, chin_err chan<- error) {
	// closing the connection and removing it from the servers slice
	conn.Close()
	conns.removeAt(i)

	// sending error to the main loop
	chin_err <- err
}

func newHealthchecker(sock common.TCPSocket) (Healthchecker, error) {
	l, err := net.Listen("tcp", sock.SockAddr.String())
	if err != nil {
		return HealthServer{}, nil
	}

	healthchecker := new(HealthServer)
	healthchecker.l = l
	healthchecker.sock = sock

	return *healthchecker, nil
}
