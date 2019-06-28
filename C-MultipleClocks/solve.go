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
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = nextInt()
	}

	rec := make([]uint64, n)
	rec[0] = uint64(t[0])
	for i := 1; i < n; i++ {
		rec[i] = lcm(rec[i-1], uint64(t[i]))
	}
	fmt.Println(rec[n-1])
}

func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b uint64) uint64 {
	return a * (b / gcd(a, b))
}
