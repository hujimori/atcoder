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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	x := nextInt()

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	sort.Sort(sort.IntSlice(slice))
	fmt.Println(slice)
	max := math.Inf(-1)

	for i := 0; i < n; i++ {
		total := x
		total -= slice[0] * i

		for j := 0; j < n-1; j++ {
			total -= slice[1] * j

			k := n - i - j
			total -= slice[2] * k
			fmt.Printf("total = %d\n", total)
			if total == 0 {
				fmt.Printf("(%d, %d, %d)\n", i, j, k)
				count := math.Max(float64(i), math.Max(float64(j), float64(k)))
				if max < count {
					max = count
				}
			}

		}
	}

	fmt.Println(max)

}
