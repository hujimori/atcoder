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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

var count = 0

func rec(slice []int, sum int, i int) bool {

	if n == sum {
		count++
		return true
	}

	if n < sum {
		return false
	}

	can := true
	if slice[i+1] == 0 {
		can = false
	} else {
		can = rec(slice, sum+1, i+1)
	}

	if slice[i+2] == 0 {
		can = false
	} else {
		can = rec(slice, sum+2, i+2)
	}

	return can
}

var n int
var slice []int

func main() {
	sc.Split(bufio.ScanWords)

	n = nextInt()
	slice = make([]int, n+1)
	slice[0] = 1

	for i := 1; i < n+1; i++ {
		slice[i] = 1
	}

	m := nextInt()
	for i := 0; i < m; i++ {
		slice[nextInt()] = 0
	}

	fmt.Println(slice)

	rec(slice, 0, 0)

	fmt.Println(count)
}
