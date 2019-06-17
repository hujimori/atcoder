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

	slice := make([]int, 3)
	for i := 0; i < 3; i++ {
		slice[i] = nextInt()
	}

	sort.Ints(slice)
	first := slice[0]
	second := slice[1]
	third := slice[2]

	// var count int
	// min := 100
	for i := second; i <= third; i++ {
		fmt.Printf("i = %d\n", i)
		count := third - i
		fmt.Println(count)

		if tmp := (third - (count+first); tmp%2 == 0 {
			count = count + tmp/2
			break
		}

	}

}
