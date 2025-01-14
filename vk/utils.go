package main

import (
	"fmt"
	"io"
)

func readMaze(in io.Reader) ([][]int, error) {
	var width, height int
	if _, err := fmt.Fscan(in, &width, &height); err != nil {
		return nil, fmt.Errorf("err when reading width/height: %v", err)
	}

	var maze = make([][]int, height)
	for i := range height {
		maze[i] = make([]int, width)
	}

	for i := range height {
		for j := range width {
			if _, err := fmt.Fscan(in, &maze[i][j]); err != nil {
				return nil, fmt.Errorf("error when reading maze element: %v", err)
			}

		}
	}

	return maze, nil
}

func readPath(in io.Reader) (coord, coord, error) {
	var start, finish coord
	if _, err := fmt.Fscan(in, &start.j, &start.i, &finish.j, &finish.i); err != nil {
		return coord{}, coord{}, fmt.Errorf("err when reading start/finish: %v", err)
	}

	return start, finish, nil
}
