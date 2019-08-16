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

	N, K := nextInt(), nextInt()
	A := make([]int, N)
	maxA := -1
	for i := 0; i < N; i++ {
		A[i] = nextInt()
		maxA = max(maxA, A[i])
	}

	gcdA := Gcd(A)

	if K <= maxA && K%gcdA == 0 {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Gcd(args []int) int {
	b := args[0]

	for _, a := range args {
		b = gcdBase(a, b)
	}
	return b
}

func gcdBase(a, b int) int {
	if b == 0 {
		return a
	}

	return gcdBase(b, a%b)
}
