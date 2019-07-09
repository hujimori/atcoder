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
	a, b, c := slice[0], slice[1], slice[2]

	cnt := 0

	if (c-b)%2 == 0 && (c-a)%2 == 0 {
		cnt += (c-b)/2 + (c-a)/2
	} else if ((c-b)%2 == 0 && (c-a)%2 == 1) || ((c-b)%2 == 1 && (c-a)%2 == 0) {
		cnt += (c-b)/2 + (c-a)/2 + 2
	} else {
		cnt += (c-b)/2 + (c-a)/2 + 1
	}
	fmt.Println(cnt)
}
