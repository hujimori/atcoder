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

	n, d := nextInt(), nextInt()

	a := d*2 + 1
	ans := 0

	if n%a == 0 {
		ans += n / a
	} else {
		ans += n/a + 1
	}

	fmt.Println(ans)
}
