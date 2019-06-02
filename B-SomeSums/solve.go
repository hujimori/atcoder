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

	N := nextInt()
	A := nextInt()
	B := nextInt()

	sum := 0

	for i := 1; i <= N; i++ {
		num := i
		digit_sum := 0
		for {
			digit_sum += num % 10
			num = num / 10

			if num == 0 {
				break
			}
		}

		if digit_sum >= A && digit_sum <= B {
			sum += i
		}
	}

	fmt.Println(sum)

}
