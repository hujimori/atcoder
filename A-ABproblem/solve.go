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

	N, A, B := nextInt(), nextInt(), nextInt()
	if A > B {
		fmt.Println(0)
	} else if N == 1 {
		if A == B {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else {
		fmt.Println((N-2)*(B-A) + 1)
	}

}
