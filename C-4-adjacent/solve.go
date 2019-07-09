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
var a []int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	if solve() {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve() bool {
	c4 := 0
	c := 0

	for i := 0; i < n; i++ {
		if a[i]%4 == 0 {
			c4++
		} else if a[i]%2 == 1 {
			c++
		}
	}

	if c4+1 == c && n == (c4+c) {
		return true
	} else if c4 < c {
		return false
	}
	return true
}
