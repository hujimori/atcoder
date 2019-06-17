package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	c := nextInt()
	k := nextInt()

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	sort.Ints(slice)

	answer := 0
	i := 0
	for i < n {
		cap := slice[i] + k
		sum := 0
		for i < n && slice[i] <= cap && sum+1 <= c {
			i++
			sum++
		}
		answer++
	}

	fmt.Println(answer)

}
