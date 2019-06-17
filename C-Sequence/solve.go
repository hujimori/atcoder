package main

import (
	"bufio"
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

	var prev int
	if slice[0] > 0 {
		prev = 1
	} else {
		prev = -1
	}

	sum := slice[0]
	for i := 1; i < n; i++ {

		if sum+slice[0] > 0 && prev > 0 {

			prev = 1
		} else if slice[i]+slice[i+1] < 0 {
			prev = -1
		}

	}

}
