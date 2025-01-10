package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numIn := flag.Int("in", 0, "specifies the number of IN-ports")
	numOut := flag.Int("out", 0, "specifies the number of OUT-ports")

	flag.Parse()

	if *numIn == 0 || *numOut == 0 {
		log.Fatal("there should be > 1 ports of each type")
	}

	var exit bool
	portMap := New(*numIn, *numOut)

	errChan := make(chan error)
	resChan := make(chan string)

	go func() {
		for {
			select {
			case err := <-errChan:
				fmt.Println(err)
			case res := <-resChan:
				fmt.Println(res)
			}
		}
	}()

	// запуск консольного клиента
	// Usage: [opeartion: read/write] [port number] [value: if writing]
	for !exit {
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')
		args := strings.Split(text[:len(text)-1], " ")

		// читаем операцию
		op := args[0]
		if op == "read" {
			if len(args) != 2 {
				fmt.Println(usage)
				continue
			}

			portNum, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			go portMap.Read(portNum, resChan, errChan)
		} else if op == "write" {
			if len(args) != 3 {
				fmt.Println(usage)
				continue
			}

			portNum, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println(err)
				continue
			}

			value, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Println(err)
				continue
			}

			boolVal, err := toBool(value)
			if err != nil {
				fmt.Println(err)
				continue
			}

			go portMap.Write(portNum, boolVal, resChan, errChan)
		} else {
			fmt.Println(usage)
		}
	}

}
