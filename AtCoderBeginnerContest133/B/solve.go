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

	n, d := nextInt(), nextInt()
	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, d)
		for j := 0; j < d; j++ {
			x[i][j] = nextInt()

		}

	}
	// fmt.Println(x)

	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			rec := 0
			for k := 0; k < d; k++ {
				rec += (x[i][k] - x[j][k]) * (x[i][k] - x[j][k])
			}
			sqrt := int(math.Sqrt(float64(rec)))
			if sqrt*sqrt == rec {
				// fmt.Println(sqrt)

				// fmt.Println(rec)
				count++
			}
		}
	}
	fmt.Println(count)

}
