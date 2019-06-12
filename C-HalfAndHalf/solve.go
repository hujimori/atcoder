package main

import (
	"bufio"
	"fmt"
	"math"
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

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	a := nextInt()
	b := nextInt()
	c := nextInt()
	x := nextInt()
	y := nextInt()

	min := math.Inf(1)

	for i := 0; i <= 100000; i++ {
		total := i*2*c + max(0, x-i)*a + max(0, y-i)*b
		if min > float64(total) {
			min = float64(total)
		}
	}

	fmt.Println(int(min))

	// if (a + b) < c*2 {
	// 	total = a*x + b*y
	// 	fmt.Println(total)
	// 	return
	// }

	// aPizzaOfseets := x
	// bPizzaOfseets := y

	// if c > y {
	// 	fmt.Println("x > y")
	// 	aPizzaOfseets = x - y
	// 	total += y * 2 * c
	// 	total += aPizzaOfseets * a
	// } else if x < y {
	// 	fmt.Println("x < y")
	// 	bPizzaOfseets = y - x
	// 	total += x * 2 * c
	// 	total += bPizzaOfseets * b
	// } else {
	// 	fmt.Println("x == y")
	// 	total += x * 2 * c
	// }

	// fmt.Println(total)

}
