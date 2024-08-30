package main

import (
	"discovery/client/client"
	"discovery/common"
	"fmt"
	"os"
)

var discoveryHost = os.Getenv("DISCOVERY_HOST")
var discoveryPort = os.Getenv("DISCOVERY_PORT")

func main() {
	if discoveryHost == "" {
		discoveryHost = "localhost"
	}

	if discoveryPort == "" {
		discoveryPort = "6970"
	}

	if len(os.Args) < 2 {
		fmt.Println("Specify handler id")
		return
	}
	handler := os.Args[1]

	client, err := client.NewClient(discoveryHost, discoveryPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Send(common.TCPMessage{Data: "pizda"}, handler)
	if err != nil {
		fmt.Println(err)
	}
}
