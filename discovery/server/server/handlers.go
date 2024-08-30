package server

import (
	"discovery/common"
	"fmt"
)

type TCPMessageHandler struct{}

func (handler *TCPMessageHandler) HandleMessage(arg common.TCPMessage, reply *common.TCPMessage) error {
	fmt.Printf("Received value: %v\n", arg)
	reply.Data = "Echo: " + arg.Data
	return nil
}

func (handler *TCPMessageHandler) HandleAnotherMessage(arg common.TCPMessage, reply *common.TCPMessage) error {
	fmt.Printf("Received another value: %v\n", arg)
	reply.Data = "rot ebal"
	return nil
}

type TCPHealthHandler struct{}

func (handler *TCPHealthHandler) HandleHealthcheck(arg common.TCPMessage, reply *common.TCPMessage) error {
	reply.Data = "ok"
	return nil
}
