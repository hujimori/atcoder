package main

import (
	"bufio"
	"fmt"
	"os"
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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[a[i]]++
		m[a[i]+1]++
		m[a[i]-1]++
	}

	ans := -1
	for _, v := range m {
		if ans < v {
			ans = v
		}
	}
	fmt.Println(ans)
}
