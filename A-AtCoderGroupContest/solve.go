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
	slice := make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		slice[i] = nextInt()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	// fmt.Println(slice)
	sum := 0

	for i := 0; i < n; i++ {
		// fmt.Println(slice[2*i+1])
		sum += slice[2*i+1]
	}

	fmt.Println(sum)
}
