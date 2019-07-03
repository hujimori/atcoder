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
	a, b := nextInt(), nextInt()

	var x int
	var y int

	extgcd(a, b, &x, &y)
	fmt.Printf("x=%d y=%d\n", x, y)
}

func extgcd(a int, b int, x *int, y *int) int {
	d := a
	if b != 0 {
		d = extgcd(b, a%b, y, x)
		*y -= (a / b) * (*x)
	} else {
		*x = 1
		*y = 0
	}
	return d
}
