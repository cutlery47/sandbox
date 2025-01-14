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
	startCoord, finishCoord, err := readPath(in)
	if err != nil {
		log.Fatal("err when reading start/finish: ", err)
	}

	// Создаем объекты-точки из координат
	start := point{coord: startCoord, weight: maze[startCoord.i][startCoord.j]}
	finish := point{coord: finishCoord, weight: maze[finishCoord.i][finishCoord.j]}

	// Решение задачи
	path := solve(maze, start, finish)

	// Запись результата в stdout
	for _, el := range path {
		fmt.Fprintf(out, "%d %d\n", el.j, el.i)
	}
	fmt.Fprintln(out, ".")
}
