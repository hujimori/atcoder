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

	N := nextInt()
	ans := 0

	for i := 1; i <= N; i++ {
		if (i >= 1 && i <= 9) || (i >= 100 && i <= 999) || (i >= 10000 && i <= 99999) {
			ans++
		}
	}

	fmt.Println(ans)
}
