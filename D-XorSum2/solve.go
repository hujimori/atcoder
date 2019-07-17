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

	sum := a[0]
	xor := a[0]
	right := 1
	res := 0
	for left := 0; left < n; left++ {
		for right < n && (sum+a[right]) == (xor^a[right]) {
			sum += a[right]
			xor ^= a[right]
			right++
		}
		res += right - left

		if right == left {
			right++
		} else {
			sum -= a[left]
			xor ^= a[left]
		}
	}

	fmt.Println(res)
}
