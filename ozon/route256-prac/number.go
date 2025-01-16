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

	res := []any{}

	var numbers int
	var number string

	fmt.Fscan(in, &numbers) // scanning the amount of values
	for range numbers {
		// running algorithm on each method
		fmt.Fscan(in, &number)
		small := calc(number)
		res = append(res, small)
	}

	for _, el := range res {
		fmt.Fprintln(out, el)
	}
}

func calc(value string) string {
	if len(value) <= 1 {
		return "0"
	}

	i := 0
	for i < len(value)-1 {
		if value[i] < value[i+1] {
			return value[:i] + value[i+1:]
		}
		i++
	}

	return value[:len(value)-1]
}
