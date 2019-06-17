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

	length := 1
	up := 0
	down := 0

	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			up = 1
		} else if slice[i-1] > slice[i] {
			down = 1
		} else {
			// i++
			continue
		}

		if up == 1 && down == 1 {
			up = 0
			down = 0

			length++
		}

	}

	fmt.Println(length)
}
