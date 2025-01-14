package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("in")
	if err != nil {
		log.Fatal("err when reading file: ", err)
	}

	// Получаем дескрипторы потоков ввода/вывода
	in := bufio.NewReader(fd)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Считываем входные данные лабиринта из stdin
	maze, err := readMaze(in)
	if err != nil {
		log.Fatal("err when reading maze: ", err)
	}

	// Считываем входные данные точек начала и конца из stdin
	start, finish, err := readPath(in)
	if err != nil {
		log.Fatal("err when reading start/finish: ", err)
	}

	// Решение задачи
	path, err := solve(maze, start, finish)
	if err != nil {
		log.Fatal("err when solving maze: ", err)
	}

	// Запись результата в stdout
	for _, el := range path {
		fmt.Fprintf(out, "%d %d\n", el.j, el.i)
	}
	fmt.Fprintln(out, ".")
}
