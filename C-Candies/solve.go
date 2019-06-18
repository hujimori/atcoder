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

var n int
var slice [][]int

func main() {
	sc.Split(bufio.ScanWords)

	slice = make([][]int, 2)

	n = nextInt()
	for i := 0; i < 2; i++ {
		slice[i] = make([]int, n)
		for j := 0; j < n; j++ {
			slice[i][j] = nextInt()
		}
	}

	// fmt.Println(slice)

	// ans := rec(3, 0, 0)

	ans := -1
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j <= i; j++ {
			sum += slice[0][j]
			// fmt.Println(slice[0][j])
		}

		sum += slice[1][i]

		for k := i + 1; k < n; k++ {
			sum += slice[1][k]
		}

		if ans < sum {
			ans = sum
		}

	}

	fmt.Println(ans)

}
