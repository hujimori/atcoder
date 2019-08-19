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

	X := nextInt()

	t := 0
	pos := 0
	for pos < X {
		t++
		pos = t * (t + 1) / 2
	}

	fmt.Println(t)

}
