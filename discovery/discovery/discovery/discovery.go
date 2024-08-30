package discovery

import (
	"discovery/common"
	"fmt"
	"net"
	"net/rpc"
)

type TCPServer struct {
	conn net.Conn
	addr net.TCPAddr
	rpc  *rpc.Client
}

type Connections []TCPServer

func (conns *Connections) printAll() {
	fmt.Println("Available connections:")
	for _, conn := range *conns {
		fmt.Printf("%v; ", conn.addr.String())
	}
	fmt.Println()
}

func (conns *Connections) removeAt(index int) {
	*conns = append((*conns)[:index], (*conns)[index+1:]...)
}

type Discoverer interface {
	Serve() error
}

type DiscoveryServer struct {
	connections   *Connections
	forwarder     Forwarder
	healthchecker Healthchecker
}

func (d DiscoveryServer) Serve() error {
	reg_err_chan := make(chan error)
	forw_err_chan := make(chan error)
	health_err_chan := make(chan error)

	go d.healthchecker.register(d.connections, health_err_chan)
	go d.healthchecker.healthcheck(d.connections, health_err_chan)
	go d.forwarder.forward(d.connections, forw_err_chan)

	for {
		select {
		case err := <-reg_err_chan:
			fmt.Println("Registration error: ", err)
		case err := <-health_err_chan:
			fmt.Println("Healthcheck error: ", err)
		case err := <-forw_err_chan:
			fmt.Println("Forwarding error: ", err)
		}
	}

}

func New(host string, regport string, forport string) (Discoverer, error) {
	healthSock, err := common.NewTCPSocket(host, regport)
	if err != nil {
		return DiscoveryServer{}, err
	}

	forwardSock, err := common.NewTCPSocket(host, forport)
	if err != nil {
		return DiscoveryServer{}, err
	}

	connections := new(Connections)

	healthchecker, err := newHealthchecker(healthSock)
	if err != nil {
		return DiscoveryServer{}, err
	}

	forwarder, err := newForwarder(forwardSock, connections)
	if err != nil {
		return DiscoveryServer{}, err
	}

	return DiscoveryServer{
		connections:   connections,
		healthchecker: healthchecker,
		forwarder:     forwarder,
	}, nil
}
