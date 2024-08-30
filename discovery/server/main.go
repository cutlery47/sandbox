package main

import (
	"discovery/server/server"
	"fmt"
	"os"
)

var discoveryHost = os.Getenv("DISCOVERY_HOST")
var discoveryPort = os.Getenv("DISCOVERY_REGISTER_PORT")

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Cli argument required")
		return
	}

	if discoveryHost == "" {
		discoveryHost = "localhost"
	}

	if discoveryPort == "" {
		discoveryPort = "6971"
	}

	serverPort := os.Args[1]
	server, err := server.New("localhost", serverPort, discoveryHost, discoveryPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

}
