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

var A, B, C int

const INF = 1000010000000000.0

func main() {
	sc.Split(bufio.ScanWords)

	A, B, C = nextInt(), nextInt(), nextInt()

	lb := 0.0
	ub := INF

	for i := 0; i < 2000000; i++ {
		mid := (ub + lb) / 2.0
		if check(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	// if !check(ub) {
	// 	ub = lb
	// }

	fmt.Printf("%.12f\n", lb)
}

func check(t float64) bool {

	f := float64(A)*t + float64(B)*math.Sin(float64(C)*t*math.Pi)
	if f < 100.0 {
		return true
	}
	return false
}
