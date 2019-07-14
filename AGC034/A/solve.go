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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	// sum := make([]uint64, n-1)
	ans := a[0]
	for i := 1; i < n; i++ {
		ans = ans ^ a[i]
	}

	if ans == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
