package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	fd, _ := os.Open("15")

	in := bufio.NewReader(fd)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var entries int
	fmt.Fscan(in, &entries)

	var numOrders int
	var numCars int
	var strArrival string

	var arrival, start *queue
	var end, capacity []int
	var val, startVal int

	for range entries {
		fmt.Fscan(in, &numOrders)
		in.ReadString('\n')

		arrival = &queue{}
		start = &queue{}

		strArrival, _ = in.ReadString('\n')
		for i, el := range strings.Split(strArrival[:len(strArrival)-1], " ") {
			val, _ = strconv.Atoi(el)
			*arrival = append(*arrival, entry{val: val, i: i})
		}

		fmt.Fscan(in, &numCars)
		end = make([]int, numCars)
		capacity = make([]int, numCars)

		for i := range numCars {
			fmt.Fscan(in, &startVal, &end[i], &capacity[i])
			*start = append(*start, entry{i: i, val: startVal})
		}

		heap.Init(arrival)
		sort.Sort(start)

		res := solve(arrival, start, end, capacity)

		for i, v := range res {
			if i > 0 {
				fmt.Fprint(out, " ")
			}
			fmt.Fprint(out, v)
		}
		fmt.Fprintln(out)
	}
}

type entry struct {
	val int
	i   int
}

type queue []entry

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Less(i, j int) bool {
	if (*q)[i].val == (*q)[j].val {
		return (*q)[i].i < (*q)[j].i
	}
	return (*q)[i].val < (*q)[j].val
}

func (q *queue) Swap(i, j int) {
	(*q)[i], (*q)[j] = (*q)[j], (*q)[i]
}

func (q *queue) Push(x any) {
	(*q) = append(*q, x.(entry))
}

func (q *queue) Pop() any {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func solve(arrival, start *queue, end, cap []int) []int {
	res := make([]int, len(*arrival))
	for i := range res {
		res[i] = -1
	}

	var ent entry

	for arrival.Len() != 0 {
		ent = heap.Pop(arrival).(entry)
		for _, car := range *start {
			if car.val <= ent.val && ent.val <= end[car.i] && cap[car.i] > 0 {
				res[ent.i] = car.i + 1
				cap[car.i]--
				break
			}
		}
	}

	return res
}
