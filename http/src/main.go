package main

import (
	"fmt"
	"xyuhttp/src/server"
)

const DefaultServerHost = "localhost"
const DefaultServerPort = 6969

func main() {
	server, err := server.NewHTTPServer(DefaultServerHost, DefaultServerPort)
	if err != nil {
		fmt.Println(err)
	}

	err = server.Serve()
	if err != nil {
		fmt.Println(err)
	}
}
