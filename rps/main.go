package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	_ "github.com/lib/pq"
)

var q = "SELECT * FROM request"
var dsn = "postgresql://postgres:12345@localhost/postgres_rps"

func main() {
	rps, err := benchmark("SELECT * FROM request", dsn, time.Millisecond*10)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("rps: ", rps)
}

func benchmark(q string, dsn string, dur time.Duration) (rps float64, err error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return rps, fmt.Errorf("Open: %v", err)
	}

	// счетчик запросов, не вернувших ошибку
	var cnt int64 = 0
	wg := &sync.WaitGroup{}

	for start := time.Now(); time.Since(start) < dur; {
		wg.Add(1)
		// параллельно рассылаем запросы в бд
		go func() {
			defer wg.Done()
			_, err := db.Query(q)
			if err != nil {
				log.Println(err)
			} else {
				// используем атомарное суммирование во избежание гонки
				atomic.AddInt64(&cnt, 1)
			}
		}()
	}

	// ожидаем завершения всех горутин
	wg.Wait()

	// приводим rps к соотношению запросы / секунда
	rps = float64(cnt) / dur.Seconds()

	return rps, nil
}
