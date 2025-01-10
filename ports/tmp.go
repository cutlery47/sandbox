package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Тип порта - булевое значение: 0 - IN, 1 - OUT
type PortType bool

const (
	In  PortType = false
	Out PortType = true
)

type Port struct {
	// Тип порта
	Type PortType
	// Значение порта
	Val bool
}

// Мапа для константного по времени поиска портов
type PortMap struct {
	// RWMutex для защиты от гонок
	// Данная программа не предусматривает такого сценария,
	// однако, если мапа будет использоваться в конкурентном окружении -
	// без блокировок не обойтись
	mu *sync.RWMutex
	// Сама мапа собственно
	mp map[int]*Port
}

// Конструктор портов
func New(in, out int) PortMap {
	num := 0
	pMap := PortMap{
		mu: &sync.RWMutex{},
		mp: map[int]*Port{},
	}

	for range in {
		v, _ := toBool(rand.Intn(2))
		port := Port{
			Type: In,
			Val:  v,
		}

		pMap.mp[num] = &port
		num++
	}

	for range out {
		v, _ := toBool(rand.Intn(2))
		port := Port{
			Type: Out,
			Val:  v,
		}

		pMap.mp[num] = &port
		num++
	}

	return pMap
}

var (
	ErrNoSuchPort     = errors.New("err: such port does not exist")
	ErrWrongPort      = errors.New("err: wrong port type")
	ErrNotConvertable = errors.New("err: cant convert value to bool")
)

// Читает значение IN-порта и возвращает его
func (pm PortMap) Read(portNum int) (bool, error) {
	// блокируем мапу на чтение
	pm.mu.RLock()
	port, ok := pm.mp[portNum]
	if !ok {
		return false, ErrNoSuchPort
	}
	pm.mu.RUnlock()

	if port.Type != In {
		return false, ErrWrongPort
	}

	return port.Val, nil
}

// Пишет новое значение OUT-порта или возвращает ошибку
// При успешной записи выводится новое значение
func (pm PortMap) Write(portNum int, val bool) error {
	port, ok := pm.mp[portNum]
	if !ok {
		return ErrNoSuchPort
	}

	if port.Type != Out {
		return ErrWrongPort
	}

	// блокируем мапу на запись
	pm.mu.Lock()
	port.Val = val
	pm.mu.Unlock()

	fmt.Println("value written:", val)
	return nil
}

func toBool(v int) (bool, error) {
	if v > 1 {
		return false, ErrNotConvertable
	}

	if v == 1 {
		return true, nil
	}
	return false, nil
}

func toInt(v bool) int {
	if v == false {
		return 0
	}
	return 1
}

const usage = "Usage: [opeartion: read/write] [port number] [value: if writing]"

func main() {
	numIn := flag.Int("in", 0, "specifies the number of IN-ports")
	numOut := flag.Int("out", 0, "specifies the number of OUT-ports")

	flag.Parse()

	if *numIn == 0 || *numOut == 0 {
		log.Fatal("there should be > 1 ports of each type")
	}

	var exit bool
	portMap := New(*numIn, *numOut)

	for !exit {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your operation: ")

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

			val, err := portMap.Read(portNum)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("value %v in port %v\n", toInt(val), portNum)
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

			if err := portMap.Write(portNum, boolVal); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println(usage)
		}
	}

}
