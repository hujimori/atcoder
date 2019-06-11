package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	k := nextInt()

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	total := 0
	for i := 0; i < k; i++ {
		total += slice[i]
	}

	fmt.Println(total)

}
