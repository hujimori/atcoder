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
	N, T := nextInt(), nextInt()
	t := make([]int, N)
	for i := 0; i < N; i++ {
		t[i] = nextInt()
	}

	// sum := make([]int, N+1)
	ans := 0

	for i := 0; i < N-1; i++ {
		if t[i+1] >= t[i]+T {
			ans += T
		} else {
			ans += t[i+1] - t[i]
		}
	}

	ans += T
	fmt.Println(ans)

}
