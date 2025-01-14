package main

import (
	"errors"
)

// Используем -1 для представления бесконечности
// т.к. golang поддерживает inf только для float64
var inf int = -1

// Структура для хранения координат лабиринта
type coord struct {
	i int
	j int
}

// Min-queue по длине пути
// Используется для эффективного нахождения точки,
// расстояние до которой минимально
type xyuQueue []xyu

type xyu struct {
	// Точка, относительно которой смотрим кратчайший путь
	cur coord
	// Длина кратчайшего пути в cur
	len int
	// Флаг, является ли точка cur начально
	start bool
}

// Мапа для хранения кратчайших путей для точек (Dijkstra)
type pathMap map[coord]path

// Структура для хранения информации о кратчайшем пути в cur
type path struct {
	// Указатель, т.к. мы хотим, чтобы при изменении в мапе изменялось значение
	// и в min-heap
	some *xyu
	// Точка, придя из которой получается кратчайший путь в cur
	prev coord
}

func solve(maze [][]int, start, finish coord) ([]coord, error) {
	if len(maze) == 0 {
		return []coord{}, nil
	}

	if maze[start.i][start.j] == 0 || maze[finish.i][finish.j] == 0 {
		return nil, errors.New("maze unsolvable")
	}

	// Хэш-сет для хранения посещенных вершин
	var visited map[coord]struct{}
	// Хэш-сет для хранения недоступных вершин
	var blocked map[coord]struct{}
	// Min-queue для хранения непомещенных вершин
	var unvisited xyuQueue
	// Мапа для хранения путей
	var paths pathMap

	visited = make(map[coord]struct{})
	unvisited, blocked = markPoints(maze, start)
	paths = createPathMap(&unvisited)

	// for len(unvisited) != 0 {

	// }

	return []coord{}, nil
}

// Создаем:
// 1) Слайс для хранения непосещенных вершин
// 2) Мапу для хранения заблокированных вершин
func markPoints(maze [][]int, start coord) (xyuQueue, map[coord]struct{}) {
	blocked := map[coord]struct{}{}
	unvisited := xyuQueue{}

	for i := range len(maze) {
		for j := range len(maze[i]) {
			if i == start.i && j == start.j {
				// Помечаем начальную вершину
				unvisited = append(unvisited, xyu{coord{i, j}, 0, true})
			} else {
				if maze[i][j] == 0 {
					// Помеаем заблокированную вершину
					blocked[coord{i, j}] = struct{}{}
				} else {
					// Помечаем непосещенную вершину
					unvisited = append(unvisited, xyu{coord{i, j}, inf, false})
				}
			}

		}
	}

	return unvisited, blocked
}

// Создаем таблицу для отслеживания кратчайших путей
func createPathMap(unvisited *xyuQueue) pathMap {
	paths := pathMap{}

	for i := range *unvisited {
		paths[(*unvisited)[i].cur] = path{some: &(*unvisited)[i], prev: coord{}}
	}

	return paths
}
