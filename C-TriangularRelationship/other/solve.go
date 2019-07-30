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

	n, k := nextInt(), nextInt()

	if k%2 == 0 {
		a := n / k
		b := (n + (k / 2)) / k
		fmt.Println(a*a*a + b*b*b)
	} else {
		a := n / k
		fmt.Println(a * a * a)
	}
}
