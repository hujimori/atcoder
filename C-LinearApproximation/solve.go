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

var N int
var A []int

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	A = make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
		A[i] -= i + 1
	}

	sort.Ints(A)
	b := A[N/2]

	ans := 0
	for i := 0; i < N; i++ {
		ans += abs(A[i] - b)
	}

	fmt.Println(ans)
	// S := N * (N + 1) / 2

	// var b int
	// if (S-sum)%N != 0 {
	// 	b = 0
	// } else {
	// 	b = (S - sum) / N
	// }

	// fmt.Println(b)

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
