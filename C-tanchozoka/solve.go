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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	res := 0
	right := 1
	for left := 0; left < n; left++ {

		for right < n && (right <= left || a[right-1] < a[right]) {
			right++
		}

		res += right - left

		if right == left {
			right++
		}
	}

	fmt.Println(res)

}
