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

	slice := make([]int, 3)
	for i := 0; i < 3; i++ {
		slice[i] = nextInt()
	}

	min := 10000
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}

			sum = slice[i] + slice[j]
			// fmt.Printf("%d + % d = %d\n", slice[i], slice[j], sum)
			if min > sum {
				min = sum
			}
		}
	}

	fmt.Println(min)
}
