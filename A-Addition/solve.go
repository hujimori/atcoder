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
	n := nextInt()
	slice := make([]int, n)

	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	odd := 0

	for i := 0; i < n; i++ {
		if slice[i]%2 == 1 {
			odd++
		}
	}

	if odd%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
