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
	// a := make([]int, n)
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[nextInt()]++
	}

	cnt := 0
	for k, v := range m {
		if k != v {
			if k > v {
				cnt += v
			} else {
				cnt += v - k
			}
		}
	}

	fmt.Println(cnt)
}
