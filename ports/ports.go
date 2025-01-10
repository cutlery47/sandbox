package main

import (
	"errors"
	"fmt"
	"math/rand"
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
func (pm PortMap) Read(portNum int, resChan chan<- string, errChan chan<- error) {
	// блокируем мапу на чтение
	pm.mu.RLock()
	port, ok := pm.mp[portNum]
	if !ok {
		errChan <- ErrNoSuchPort
		return
	}
	pm.mu.RUnlock()

	if port.Type != In {
		errChan <- ErrNoSuchPort
		return
	}

	resChan <- fmt.Sprintf("read port value: %v", port.Val)
}

// Пишет новое значение OUT-порта или возвращает ошибку
// При успешной записи выводится новое значение
func (pm PortMap) Write(portNum int, val bool, resChan chan<- string, errChan chan<- error) {
	port, ok := pm.mp[portNum]
	if !ok {
		errChan <- ErrNoSuchPort
		return
	}

	if port.Type != Out {
		errChan <- ErrNoSuchPort
		return
	}

	// блокируем мапу на запись
	pm.mu.Lock()
	port.Val = val
	pm.mu.Unlock()

	resChan <- fmt.Sprintf("port %v value written: %v", portNum, toInt(val))
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
