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

	n := nextInt()

	can := false
	var total int

L:
	for i := 0; i <= 100/7; i++ {
		for j := 0; j <= 100/4; j++ {
			total = i*4 + j*7

			if total == n {
				fmt.Printf("%d = %d*4 + %d*7\n", total, i, j)
				can = true
				break L
			}
		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
