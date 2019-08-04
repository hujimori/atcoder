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

	A, B, C := nextInt(), nextInt(), nextInt()

	ans := C - (A - B)

	if ans <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
