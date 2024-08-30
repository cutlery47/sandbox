package main

import (
	"discovery/discovery/discovery"
	"fmt"
	"os"
)

var Host = os.Getenv("DISCOVERY_HOST")
var RegisterPort = os.Getenv("DISCOVERY_REGISTER_PORT")
var ForwardingPort = os.Getenv("DISCOVERY_FORWARDING_PORT")

func main() {
	if Host == "" {
		Host = "localhost"
	}

	if RegisterPort == "" {
		RegisterPort = "6971"
	}

	if ForwardingPort == "" {
		ForwardingPort = "6970"
	}

	discovery, err := discovery.New(Host, RegisterPort, ForwardingPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = discovery.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
}
