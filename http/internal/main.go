package main

import (
	"fmt"
	"xyuhttp/internal/server"

	"github.com/pkg/errors"
)

const DefaultServerHost = "localhost"
const DefaultServerPort = 6969

func main() {
	server, err := server.NewHTTPServer(DefaultServerHost, DefaultServerPort)
	if err != nil {
		fmt.Println(err)
		errors.New("xyu")
	}

	err = server.Serve()
	if err != nil {
		fmt.Println(err)
	}
}
