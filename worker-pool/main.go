package main

import (
	"log"
	"math/rand"
	"time"
)

type Work struct {
	val int
}

type Result struct {
	res int
}

type Worker struct {
	id       int
	workChan <-chan Work
	resChan  chan<- Result
}

func InitWorker(id int, workChan <-chan Work, resChan chan<- Result) Worker {
	return Worker{id: id, workChan: workChan, resChan: resChan}
}

func (w Worker) Work() {
	for {
		work := <-w.workChan
		log.Printf("worker %v has received the work\n", w.id)
		res := w.CalcFact(work.val)
		w.resChan <- Result{res: res}
		log.Printf("worker %v has sent the result\n", w.id)
	}
}

func (w Worker) CalcFact(num int) int {
	if num == 0 || num == 1 {
		return num
	} else {
		return num * w.CalcFact(num-1)
	}
}

func CreateWorkers(num int) (chan<- Work, <-chan Result) {
	workChan := make(chan Work)
	resChan := make(chan Result)

	for v := range num {
		worker := InitWorker(v, workChan, resChan)
		go worker.Work()
	}

	return workChan, resChan
}

func main() {
	workers := 10

	workChan, resChan := CreateWorkers(workers)

	go func() {
		for {
			log.Println(<-resChan)
		}
	}()

	go func() {
		for {
			val := rand.Intn(10)
			workChan <- Work{val: val}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {

	}

}
