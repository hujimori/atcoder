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
	x := nextInt()

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	d := 0
	count := 1

	for i := 1; i < n+1; i++ {
		d = d + slice[i-1]
		// fmt.Println(d)
		if d <= x {
			count++
		}
	}

	fmt.Println(count)
}
