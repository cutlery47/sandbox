package main

import (
	"container/heap"
	"errors"
	"math"
)

// Используем MaxInt для представления бесконечности
// т.к. го поддерживает inf только для float64
var inf int = math.MaxInt

// Структура для хранения координат лабиринта
type coord struct {
	i int
	j int
}

type xyu struct {
	// Точка, относительно которой смотрим кратчайший путь
	cur coord
	// Длина кратчайшего пути в cur
	len int
	// Флаг, является ли точка cur начальной
	start bool
}

// Min-queue по длине пути
// Используется для эффективного нахождения точки,
// расстояние до которой минимально
type xyuQueue []xyu

// Имллементация heap.Interface
func (q *xyuQueue) Len() int {
	return len(*q)
}

func (q *xyuQueue) Less(i, j int) bool {
	if (*q)[i].len < (*q)[j].len {
		return true
	}
	return false
}

func (q *xyuQueue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *xyuQueue) Push(x any) {
	*q = append(*q, x.(xyu))
}

func (q *xyuQueue) Pop() any {
	val := (*q)[q.Len()-1]
	(*q) = (*q)[:q.Len()-1]
	return val
}

func (q *xyuQueue) Update(upd coord, newLen int) {
	for i := range *q {
		if (*q)[i].cur == upd {
			(*q)[i].len = newLen
			break
		}
	}
}

// Структура для хранения информации о кратчайшем пути в cur
type path struct {
	// Указатель, т.к. мы хотим, чтобы при изменении в мапе изменялось значение
	// и в min-heap
	len int
	// Точка, придя из которой получается кратчайший путь в cur
	prev coord
}

// Мапа для хранения кратчайших путей для точек (Dijkstra)
type pathMap map[coord]path

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
	var unvisited *xyuQueue
	// Мапа для хранения путей
	var paths pathMap

	visited = make(map[coord]struct{})
	unvisited, blocked = markPoints(maze, start)
	paths = createPathMap(unvisited)

	// Переменная для хранения текущей вершины
	var node xyu
	// Переменная для хранения веса текущей вершины
	var nodeWeight int
	// Слайс вершин, до которых можно дотянуться из node
	var reachable []coord
	// Переменная для хранения координат вершин-кандидатов reachable
	var candidate coord
	// Переменные для хранения значений pathMap
	var nodePath, curPath path

	for unvisited.Len() != 0 {
		curPath = paths[node.cur]

		// Достаем вершину с минимальным расстоянием и ее вес
		node = heap.Pop(unvisited).(xyu)
		nodeWeight = maze[node.cur.i][node.cur.j]

		// Добавляем вершину в reachable только если она досягаема, не заблокирована и не посещена

		if node.cur.i != 0 {
			candidate = coord{node.cur.i - 1, node.cur.j}
			if _, ok := blocked[candidate]; !ok {
				if _, ok := visited[candidate]; !ok {
					reachable = append(reachable, candidate)
				}
			}
		}

		if node.cur.i != len(maze)-1 {
			candidate = coord{node.cur.i + 1, node.cur.j}
			if _, ok := blocked[candidate]; !ok {
				if _, ok := visited[candidate]; !ok {
					reachable = append(reachable, candidate)
				}
			}
		}

		if node.cur.j != 0 {
			candidate = coord{node.cur.i, node.cur.j - 1}
			if _, ok := blocked[candidate]; !ok {
				if _, ok := visited[candidate]; !ok {
					reachable = append(reachable, candidate)
				}
			}
		}

		if node.cur.j != len(maze[node.cur.i])-1 {
			candidate = coord{node.cur.i, node.cur.j + 1}
			if _, ok := blocked[candidate]; !ok {
				if _, ok := visited[candidate]; !ok {
					reachable = append(reachable, candidate)
				}
			}
		}

		// Обрабатываем досягаемые вершины
		for _, el := range reachable {
			// Получаем из мапы путь до текущей вершины
			curPath = paths[node.cur]
			// Получаем из мапы путь до досягаемой вершины
			nodePath = paths[el]

			// Обновляем путь до досягаемой вершины, если нашли более короткий
			if nodePath.len == inf || curPath.len+nodeWeight < nodePath.len {
				// Обновляем кратчайший путь до достижимой вершины и обновляем ее предшественника
				nodePath.len = curPath.len + nodeWeight
				nodePath.prev = node.cur

				// Обновляем запись в мапе путей
				paths[el] = nodePath
				// Обновляем запись в очереди вершин
				unvisited.Update(el, nodePath.len)
			}
		}

		// Помечаем текущую вершину посещенной
		visited[node.cur] = struct{}{}
		// Heapify, т.к. значения в очереди могли измениться при обновлении путей
		heap.Init(unvisited)

		// Подчищаемся
		reachable = []coord{}
		candidate = coord{}
		nodePath, curPath = path{}, path{}
		node = xyu{}
	}

	return restorePath(paths, finish), nil
}

// Создаем:
// 1) Слайс для хранения непосещенных вершин
// 2) Мапу для хранения заблокированных вершин
func markPoints(maze [][]int, start coord) (*xyuQueue, map[coord]struct{}) {
	blocked := map[coord]struct{}{}
	unvisited := &xyuQueue{}
	heap.Init(unvisited)

	for i := range len(maze) {
		for j := range len(maze[i]) {
			if i == start.i && j == start.j {
				// Помечаем начальную вершину
				heap.Push(unvisited, xyu{coord{i, j}, 0, true})
			} else {
				if maze[i][j] == 0 {
					// Помеаем заблокированную вершину
					blocked[coord{i, j}] = struct{}{}
				} else {
					// Помечаем непосещенную вершину
					heap.Push(unvisited, xyu{coord{i, j}, inf, false})
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
		paths[(*unvisited)[i].cur] = path{len: (*unvisited)[i].len, prev: coord{}}
	}

	return paths
}

// Восстанавливаем путь по мапе
func restorePath(paths pathMap, finish coord) []coord {
	fullPath := []coord{finish}

	// 	Бежим с конца, пока не дойдем до старта
	for curPath := paths[finish]; curPath.len != 0; curPath = paths[curPath.prev] {
		fullPath = append(fullPath, curPath.prev)
	}

	// 	Переворачиваем получившийся путь
	for i := range len(fullPath) / 2 {
		fullPath[i], fullPath[len(fullPath)-1-i] = fullPath[len(fullPath)-1-i], fullPath[i]
	}

	return fullPath
}
