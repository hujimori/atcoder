package main

import (
	"bufio"
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

var n int
var L []int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	L = make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = nextInt()
	}
	solve()
}

func solve() {
	ans := 0

	for n > 1 {
		// 一番短いmii1、次に短いmii2を求める
		mii1, mii2 := 0, 1
		if L[mii1] > L[mii2] {
			swap(&mii1, &mii2)
		}

		for i := 2; i < n; i++ {
			if L[i] < L[mii1] {
				mii2 = mii1
				mii1 = i
			} else if L[i] < L[mii2] {
				mii2 = i
			}
		}

		// それらを併合
		t := L[mii1] + L[mii2]
		ans += t

		if mii1 == n-1 {
			swap(&mii1, &mii2)
		}
		L[mii1] = t
		L[mii2] = L[n-1]
		n--
	}

	println(ans)
}

func swap(a, b *int) {
	tmp := a
	a = b
	b = tmp
}
