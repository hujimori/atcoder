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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	N, X := nextInt(), nextInt()
	x := make([]int, N+1)
	x[0] = X
	for i := 1; i < N+1; i++ {
		x[i] = nextInt()
	}
	sort.Ints(x)
	fmt.Println(x)
	// min := 1000000000000
	n := make([]int, N+1)
	for i := 0; i < N+1; i++ {
		n[i] = abs(x[i+1] - x[i])
	}

	res := make([]int, len(n))
	res[0] = n[1]
	for i := 1; i < len(n)-1; i++ {
		gcd(n[i])

	}

	fmt.Println(n)

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
