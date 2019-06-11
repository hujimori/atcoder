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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	var a int
	var b int
	min := math.Inf(1)

	for i := 1; i < n; i++ {
		a = i
		b = n - i
		// fmt.Printf("N = %d, a = %d, b = %d\n", n, a, b)

		aDigitSum := 0
		bDigitSum := 0

		for {
			aDigitSum += a % 10
			a /= 10

			if a == 0 {
				break
			}
		}

		for {
			bDigitSum += b % 10
			b /= 10

			if b == 0 {
				break
			}
		}

		total := float64(aDigitSum + bDigitSum)

		// fmt.Printf("total = %f, A = %d, B = %d\n", total, aDigitSum, bDigitSum)

		if min > total {
			// fmt.Printf("N = %d, a = %d, b = %d\n", i, a, b)
			min = total
			// fmt.Println(min)

		}
	}

	fmt.Println(min)
}
