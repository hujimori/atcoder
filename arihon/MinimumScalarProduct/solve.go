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
	v1 := make([]int, n)
	v2 := make([]int, n)

	for i := 0; i < n; i++ {
		v1[i], v2[i] = nextInt(), nextInt()
	}

	sort.Ints(v1)
	sort.Sort(sort.Reverse(sort.IntSlice(v2)))
	fmt.Println(v1)
	fmt.Println(v2)

	ans := int64(0)
	for i := 0; i < n; i++ {
		ans += int64(v1[i]) * int64(v2[i])
	}
	println(ans)
}
