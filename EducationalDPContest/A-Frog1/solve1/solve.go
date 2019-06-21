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

var slice []int

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	slice = make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	fmt.Println(solve(n-1, 0))
}

func solve(i, sum int) int {

	if i == 0 {
		return sum
	}

	ret1 := 1000000000
	if i-1 >= 0 {
		ret1 = solve(i-1, sum+abs(slice[i]-slice[i-1]))
	}

	ret2 := 1000000000
	if i-2 >= 0 {
		ret2 = solve(i-2, sum+abs(slice[i]-slice[i-2]))
	}

	return min(ret1, ret2)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
