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

func main() {
	sc.Split(bufio.ScanWords)

	k := nextInt()
	s := nextInt()
	sum := 0
	count := 0
	for i := 0; i <= k; i++ {
		sum = s - i
		for j := 0; j <= int(math.Min(float64(sum), float64(k))); j++ {

			if (sum - j) <= k {
				// fmt.Printf("X = %d, Y = %d, Z = %d\n", i, j, sum-j)
				count++
			}
		}
	}

	fmt.Println(count)
}
