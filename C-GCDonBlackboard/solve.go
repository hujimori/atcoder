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
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	max := -1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			res := gcd(slice[i], slice[j])

			if max < res {
				max = res
			}
		}
	}
	fmt.Println(max)
}

func gcd(a, b int) int {

	if a < b {
		tmp := a
		a = b
		b = tmp
	}

	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
