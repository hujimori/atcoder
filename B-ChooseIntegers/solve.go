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

	a := nextInt()
	b := nextInt()
	c := nextInt()

	can := false

	sum := 0
	for i := 1; ; i++ {
		sum += a
		// fmt.Println(sum)
		if sum%b == c {
			can = true
		}

		if sum%b == 0 {
			break
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
