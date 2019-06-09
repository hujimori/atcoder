package main

import (
	"bufio"
	"fmt"
	"math"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	var s1 int
	var s2 int
	min := math.Inf(1)
	for t := 1; t <= n; t++ {
		s1 = 0
		s2 = 0
		for i, w := range slice {
			if i <= t {
				s1 += w
			}
			if t < i {
				s2 += w
			}
		}
		abs := math.Abs(float64(s1) - float64(s2))
		if min > abs {
			min = abs
		}
	}

	fmt.Println(int64(min))
}
