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

	L, R := nextInt(), nextInt()

	if R-L > 2019 {
		fmt.Println(0)
		return
	}

	min := 100000000

	for i := L; i <= R-1; i++ {
		for j := L + 1; j <= R; j++ {
			if min > i*j%2019 {
				min = i * j % 2019
			}
		}
	}

	fmt.Println(min)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
