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

	px := make([]int, 2)
	py := make([]int, 2)
	for i := 0; i < 2; i++ {
		px[i], py[i] = nextInt(), nextInt()
	}

	ans := gcd(abs(px[0]-px[1]), abs(py[0]-py[1])) - 1
	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
