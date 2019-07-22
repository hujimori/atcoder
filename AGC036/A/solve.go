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

	s := nextInt()

	var a, b int
	sup := 1000000000

	if s%sup == 0 {
		a = s / sup
		b = 0
	} else {
		a = s/sup + 1
		b = sup - s%sup
	}

	fmt.Printf("0 0 %d 1 %d %d\n", sup, b, a)

}

// func main() {
// 	sc.Split(bufio.ScanWords)

// 	s := nextInt()

// 	x1, x2, x3 := 0, 0, 0
// 	y1, y2, y3 := 0, 0, 0

// 	var y23 int
// 	for i := 1; ; i++ {
// 		if s%i == 0 && s/i >= 2 {
// 			y23 = s / i
// 			x1 = i
// 			break
// 		}
// 	}

// 	y3 = 0
// 	y2 = y23 + y3
// 	y1 = y2 + 100

// 	fmt.Printf("%d %d %d %d %d %d\n", x1, y1, x2, y2, x3, y3)

// }
