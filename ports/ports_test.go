package main

import (
	"testing"
	"time"
)

var (
	inPorts, outPorts = 5, 5
	portMap           = New(inPorts, outPorts)
	errChan, resChan  = make(chan error), make(chan string)
)

// попытка чтения из несуществующего порта
func TestReadNonexistantPort(t *testing.T) {
	go portMap.Read(100, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != ErrNoSuchPort {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Fatal("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}

// попытка записис в несузествующий порт
func TestWriteNonexistantPort(t *testing.T) {
	go portMap.Write(100, false, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != ErrNoSuchPort {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Fatal("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}

// попытка чтения из OUT-порта
func TestReadWrongType(t *testing.T) {
	go portMap.Read(5, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != ErrWrongPort {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Fatal("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}

// попытка записи в IN-порт
func TestWriteWrongType(t *testing.T) {
	go portMap.Write(0, false, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != ErrWrongPort {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Fatal("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}

func TestCorrectRead(t *testing.T) {
	go portMap.Read(0, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Log("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}

func TestCorrectWrite(t *testing.T) {
	go portMap.Write(8, true, resChan, errChan)
	time.Sleep(250 * time.Millisecond)

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatal("error: ", err)
		}
	case res := <-resChan:
		t.Log("res: ", res)
	default:
		t.Fatal("didn't receive response")
	}
}
