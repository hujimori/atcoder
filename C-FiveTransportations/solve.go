package main

import (
	"bufio"
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
	slice := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		slice[i] = nextInt()
	}

	city := make([]int, n+1)
	city[0] = n

	for i := 1; i < len(city); i++ {
		for city[i] <= slice[i] {

		}

	}

}
