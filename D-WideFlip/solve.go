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

const INF = 10000000000000000

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	ans := INF
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			ans = min(ans, max(i+1, n-1-i))
		}
	}

	if ans == INF {
		ans = n
	}

	fmt.Println(ans)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
