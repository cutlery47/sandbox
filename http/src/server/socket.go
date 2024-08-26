package server

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

type TCPServerSocket struct {
	SockAddr net.TCPAddr
}

func newTCPServerSocket(host string, port int) (TCPServerSocket, error) {
	if port < 0 || port > (2<<16-1) {
		return TCPServerSocket{}, errors.New("socket port must be in [0, 2^16]")
	}

	ip, err := stringHostToIPv4(host)
	if err != nil {
		return TCPServerSocket{}, err
	}

	sock_addr := net.TCPAddr{
		Port: port,
		IP:   ip,
	}

	return TCPServerSocket{SockAddr: sock_addr}, nil
}

func stringHostToIPv4(host string) (net.IP, error) {
	if host == "localhost" {
		return net.IP{127, 0, 0, 1}, nil
	}

	str_bytes := strings.Split(host, ".")
	if len(str_bytes) != 4 {
		return nil, errors.New("socket host must be in IPv4 format")
	}

	IP := net.IP{}

	for i, val := range str_bytes {
		ip_byte, err := strconv.Atoi(val)

		if ip_byte < 0 || ip_byte > (2<<8) {
			return nil, errors.New("each IPv4 segment must be a byte")
		}

		if err != nil {
			return nil, errors.New(err.Error())
		}

		IP[i] = byte(ip_byte)
	}

	return IP, nil
}
