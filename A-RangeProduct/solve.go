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

	a, b := nextInt(), nextInt()
	if a <= 0 && b >= 0 {
		fmt.Println("Zero")
		return
	}

	if a > 0 {
		fmt.Println("Positive")
		return
	}

	if b < 0 {
		cnt := b - a + 1
		if cnt%2 == 0 {
			fmt.Println("Positive")
		} else {
			fmt.Println("Negative")
		}
	}
}
