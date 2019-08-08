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

var X, Y int

func main() {
	sc.Split(bufio.ScanWords)
	X, Y = nextInt(), nextInt()

	if abs(X-Y) <= 1 {
		fmt.Println("Brown")
	} else {
		fmt.Println("Alice")
	}

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
