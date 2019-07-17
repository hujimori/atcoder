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

	m := map[int]int{}

	res := 0
	right := 0
	for left := 0; left < n; left++ {

		for right < n && m[a[right]] < 1 {
			m[a[right]]++
			right++
		}

		res = max(res, right-left)
		if right == left {
			right++
		} else {
			m[a[left]]--
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
