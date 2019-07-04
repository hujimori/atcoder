package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var n int
var a []int
var k int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	sort.Ints(a)

	k = nextInt()

	fmt.Println(lowerBound())
}

func lowerBound() int {
	// 解の存在範囲を初期化
	lb := -1
	ub := n
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if a[mid] >= k {
			// midが条件を満たせば、解の存在範囲は(lb, mid)
			ub = mid
		} else {
			// midが条件を満たさなければ、解の存在範囲は(mid, ub)
			lb = mid
		}
	}
	return ub
}
