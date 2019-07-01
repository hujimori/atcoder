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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = nextInt()
	}
	sort.Ints(d)

	// fmt.Println(d)

	fmt.Println(d[n/2] - d[n/2-1])
}
