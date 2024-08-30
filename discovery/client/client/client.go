package client

import (
	"discovery/common"
	"errors"
	"fmt"
	"net/rpc"
)

type Client struct {
	DiscoveryAddr string
	DiscoveryPort string
	RpcClient     *rpc.Client
}

func (client *Client) Send(msg common.TCPMessage, handler string) error {
	// reply message buffer
	reply := common.TCPMessage{}

	switch handler {
	case "1":
		err := client.RpcClient.Call("TCPMessageHandler.HandleMessage", msg, &reply)
		if err != nil {
			return err
		}
	case "2":
		err := client.RpcClient.Call("TCPMessageHandler.HandleAnotherMessage", msg, &reply)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported handler id")
	}

	fmt.Printf("Response: %v\n", reply)

	return nil
}

func NewClient(addr string, port string) (Client, error) {
	rpc, err := rpc.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return Client{}, err
	}

	client := Client{
		DiscoveryAddr: addr,
		DiscoveryPort: port,
		RpcClient:     rpc,
	}

	return client, nil
}
