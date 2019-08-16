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

	N := nextInt()
	L := make([]int, 2*N)
	for i := 0; i < 2*N; i++ {
		L[i] = nextInt()
	}

	sort.Ints(L)
	ans := 0
	for i := 0; i < 2*N; i += 2 {
		ans += L[i]
	}

	fmt.Println(ans)
}
