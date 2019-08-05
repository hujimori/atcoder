package main

import (
	"bufio"
	"fmt"
	"math"
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

const eps = 1e-10

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()
	x := make([]float64, N)
	y := make([]float64, N)
	for i := 0; i < N; i++ {
		x[i], y[i] = float64(nextInt()), float64(nextInt())
	}

	total := N * (N - 1) * (N - 2) / 6
	cnt90 := 0
	cnt := 0

	pi := math.Acos(-1)
	for i := 0; i < N; i++ {
		var v []float64
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			v = append(v, math.Atan2(y[j]-y[i], x[j]-x[i]))
		}

		sort.Float64s(v)
		for j := 0; j < N-1; j++ {
			v = append(v, v[j]+2*pi)
		}
		// fmt.Println(v)

		for j := 0; j < N-1; j++ {
			cnt90 += UpperBound(v, v[j]+pi/2+eps) - LowerBound(v, v[j]+pi/2-eps)
			cnt += LowerBound(v, v[j]+pi) - UpperBound(v, v[j]+pi/2+eps)
		}
	}
	// fmt.Println(cnt90)
	fmt.Printf("%d %d %d\n", total-cnt-cnt90, cnt90, cnt)

}

func LowerBound(t []float64, k float64) int {
	lb := -1
	ub := len(t)
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] >= k {
			ub = mid
		} else {
			lb = mid
		}
	}
	return ub
}

func UpperBound(t []float64, k float64) int {
	lb := -1
	ub := len(t)
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] <= k {
			lb = mid
		} else {
			ub = mid
		}
	}
	return ub
}
