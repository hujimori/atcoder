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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	card := map[string]int{}

	for i := 0; i < n; i++ {
		str := nextLine()
		if _, ok := card[str]; ok {
			card[str]++
		} else {
			card[str] = 1
		}
	}

	m := nextInt()

	for i := 0; i < m; i++ {
		str := nextLine()
		if _, ok := card[str]; ok {
			card[str]--
		} else {
			card[str] = -1
		}
	}

	// fmt.Println(card)

	max := math.Inf(-1)
	for _, v := range card {

		if max < float64(v) {
			max = float64(v)
		}
	}

	if max < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(max)
	}

}
