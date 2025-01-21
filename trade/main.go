package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const inf = math.MaxInt

/*
задача сводится к нахождению отрицательных циклов в графе
алгоритм бэллмана-форда отлично бы подошел для этой задачи,
однако он под запретом (что печально ;d)

было решено реализовать алгоритм на основе обычного обхода графа в глубину
суть алгоритма следующая:

1. собираем матрицу смежности из входных данных

2. из КАЖДОЙ вершины запускаем обход в глубину (почему из каждой см. ниже)

 3. идем до тех пор, пока не попали в вершину, в которой уже ранее бывали

 4. если раньше количество монет на этой вершине было меньше - значит мы где-то
    смогли обменять предметы на большую сумму и при этом сохранили изначальный предмет
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var items, trades int
	fmt.Fscanln(in, &items, &trades)

	// инициализируем матрицу смежности
	adj := make([][]int, items)
	for i := range adj {
		adj[i] = make([]int, items)
	}

	// по умолчанию, заполяем матрицу смежности бесконечностями
	for i := range len(adj) {
		for j := range len(adj[i]) {
			adj[i][j] = inf
		}
	}

	var u, v, w int
	// заполняем матрицу смежности значениями из stdin
	for range trades {
		fmt.Fscanln(in, &u, &v, &w)
		adj[u-1][v-1] = w
	}

	visited := make(map[int]int)

	// пробегаем граф, начиная со всех возможных вершин, т.к.
	// граф может иметь несколько компонент связности
	// (не из любой вершины можно попасть в любую другую)
	for i := range items {
		if path, ok := traverse(i, 0, []int{}, visited, &adj); !ok {
			// нашли читерский путь
			fmt.Fprintln(out, "YES")

			cycle := detect(path)
			for i, el := range cycle {
				fmt.Fprintf(out, "%v", el)

				if i != len(cycle)-1 {
					fmt.Fprint(out, " ")
				} else {
					fmt.Fprintln(out, "")
				}
			}
			return
		}
	}

	// читерский путь не был найден
	fmt.Fprintln(out, "NO")
}

// обход графа в глубину
func traverse(cur, curMoney int, curPath []int, visited map[int]int, adj *[][]int) ([]int, bool) {
	// добавляем текущую вершину в слайс, хранящий пройденный путь
	curPath = append(curPath, cur+1)

	if prevMoney, ok := visited[cur]; ok {
		// прибыли в ранее посещенную вершину с большим кол-вом монет
		if curMoney > prevMoney {
			return curPath, false
		}
		return curPath, true
	}

	// бежим по всем вершинам
	for j := range (*adj)[cur] {
		// пропускаем вершину, если она текущая и не имеет ребра в себя или недосягаема
		if (cur == j && (*adj)[cur][j] == 0) || (*adj)[cur][j] == inf {
			continue
		}

		// копируем мапу посещенных вершин
		newVisited := make(map[int]int)
		copyMap(newVisited, visited)
		newVisited[cur] = curMoney

		// копируем слайс с текущим путем
		newPath := make([]int, len(curPath))
		copy(newPath, curPath)

		// вычитаем вес ребра, по которому идем в смежную вершину
		newMoney := curMoney - (*adj)[cur][j]

		// рекурсивно обходим весь граф
		if path, ok := traverse(j, newMoney, newPath, newVisited, adj); !ok {
			return path, false
		}
	}

	return curPath, true
}

/*
вычлениям цикл из проблемного маршрута

функция имеет место, потому что может быть случай,
когда мы переходим в вершину, являющуюся частью плохого цикла, по нулевому ребру.
в таком случае, алгоритм обнаружит путь с циклом, однако положит в него вершину
из которой мы пришли в этот цикл, а она лишняя...
*/
func detect(cycle []int) []int {
	var i int
	for cycle[i] != cycle[len(cycle)-1] {
		i++
	}

	return cycle[i:]
}

func copyMap(dst, src map[int]int) {
	for k, v := range src {
		dst[k] = v
	}
}
