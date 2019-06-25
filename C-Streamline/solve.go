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
	n, m := nextInt(), nextInt()

	x := make([]int, m)
	for i := 0; i < m; i++ {
		x[i] = nextInt()
	}
	sort.Ints(x)

	if n >= m {
		fmt.Println(0)
		return
	}

	slice := make([]int, m)
	ans := 0
	for i := 1; i < m; i++ {
		slice[i] = x[i] - x[i-1]
		ans += slice[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	for i := 0; i < n-1; i++ {
		ans -= slice[i]
	}

	fmt.Println(ans)
}
