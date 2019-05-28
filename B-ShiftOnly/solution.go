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
	var n int
	fmt.Scan(&n)
	A := make([]int, n, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	loop := true
	count := 0
	for loop {

		for i := 0; i < n; i++ {
			if A[i]%2 == 1 {
				loop = false
			}
			A[i] = A[i] / 2
		}
		if loop {
			count++
		}

	}

	fmt.Println(count)
}
