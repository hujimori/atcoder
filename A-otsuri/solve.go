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
	// sc.Split(bufio.ScanWords)
	var input int
	fmt.Scan(&input)

	sum := 1000 - input
	coin := [6]int{500, 100, 50, 10, 5, 1}
	count := 0
	for i := 0; i < len(coin); i++ {
		c := sum / coin[i]
		sum -= coin[i] * c
		count += c

	}

	fmt.Println(count)
}
