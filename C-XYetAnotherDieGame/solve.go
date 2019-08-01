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

	x := nextInt()

	ans := x / 11 * 2

	if x%11 > 0 {
		if x%11 > 6 {
			ans += 2
		} else {
			ans++
		}
	}

	fmt.Println(ans)
}
