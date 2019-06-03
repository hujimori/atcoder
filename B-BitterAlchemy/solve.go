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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	x := nextInt()

	sum := x
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		m := nextInt()
		slice[i] = m
		sum -= m
	}

	sort.Ints(slice)
	// fmt.Println(slice)
	count := n
	// fmt.Println(count)
	for sum > 0 {
		sum -= slice[0]
		if sum < 0 {
			break
		}

		count++

		// fmt.Printf("sum = %d\n", sum)
	}

	fmt.Println(count)
}
