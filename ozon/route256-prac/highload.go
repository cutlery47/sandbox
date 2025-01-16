package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tasks int
	var seq string

	// scan tasks
	fmt.Fscan(in, &tasks)

	for range tasks {
		fmt.Fscan(in, &seq)
		if solve(seq) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}

func solve(seq string) bool {
	if seq[len(seq)-1] != 'D' || seq[0] != 'M' {
		return false
	}

	stack := []rune{}
	for _, el := range seq {
		switch el {
		case 'M':
			if len(stack) != 0 {
				return false
			}
			stack = append(stack, el)
		case 'D':
			if len(stack) == 0 {
				return false
			}
			stack = []rune{}
		case 'R':
			if len(stack) == 0 || stack[len(stack)-1] != 'M' {
				return false
			}
			stack = append(stack, el)
		case 'C':
			if len(stack) == 0 {
				return false
			}
			stack = []rune{}
		}
	}

	return true
}
