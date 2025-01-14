package main

type point struct {
	coord
	weight int
}

type coord struct {
	i int
	j int
}

func solve(maze [][]int, start, finish point) []coord {
	if len(maze) == 0 {
		return []coord{}
	}

	return traverse(maze, 0, []coord{}, start, finish)
}

func traverse(maze [][]int, curWeight int, curPath []coord, cur, finish point) []coord {
	return []coord{}
}
