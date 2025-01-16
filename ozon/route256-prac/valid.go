package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var batches int
	fmt.Fscan(in, &batches)

	var amount int

	for range batches {
		fmt.Fscan(in, &amount)
		in.ReadString('\n')

		input := []int{}
		strInput, _ := in.ReadString('\n')

		for _, el := range strings.Split(strInput[:len(strInput)-1], " ") {
			val, _ := strconv.Atoi(el)
			input = append(input, val)
		}

		strAnswer := sortToString(input)

		strOutput, _ := in.ReadString('\n')

		if strAnswer != strOutput[:len(strOutput)-1] {
			out.WriteString("NO\n")
		} else {
			out.WriteString("YES\n")
		}

	}

}

func sortToString(input []int) string {
	var res strings.Builder
	slices.Sort(input)

	for _, el := range input {
		res.WriteString(fmt.Sprintf("%v ", el))
	}

	return res.String()[:res.Len()-1]
}
