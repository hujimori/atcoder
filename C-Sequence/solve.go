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

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	// sum := make([]int, n+1)
	// for i := 0; i < n; i++ {
	// 	sum[i+1] = sum[i] + slice[i]
	// }

	// cnt1 := 0
	// tmp1 := sum[1:]
	// for i := 0; i < n; i++ {
	// 	if i%2 == 0 {
	// 		if tmp1[i] <= 0 {
	// 			cnt1 += abs(1 - tmp1[i])
	// 			tmp1[i] = 1
	// 		}
	// 	} else {
	// 		if tmp1[i] >= 0 {
	// 			cnt1 += abs(-1 - tmp1[i])
	// 			tmp1[i] = -1
	// 		}
	// 	}
	// }

	// fmt.Println(cnt1)

	// cnt2 := 0
	// tmp2 := sum[1:]
	// for i := 0; i < n; i++ {
	// 	if i%2 == 0 {
	// 		if tmp2[i] >= 0 {
	// 			cnt2 += abs(-1 - tmp2[i])
	// 			tmp2[i] = -1
	// 		}
	// 	} else {
	// 		if tmp2[i] <= 0 {
	// 			cnt2 += abs(1 - tmp2[i])
	// 			tmp2[i] = 1
	// 		}
	// 	}
	// }

	// fmt.Println(cnt2)

	ans1 := solve(slice, 1)
	ans2 := solve(slice, 0)
	fmt.Println(min(ans1, ans2))
}

func solve(a []int, r int) int {
	cnt := 0
	sum := 0

	for i := 0; i < len(a); i++ {
		if i%2 == r {
			if a[i]+sum <= 0 {
				cnt += abs(1 - a[i] - sum)
				sum = 1
			} else {
				sum += a[i]
			}
		} else {
			if a[i]+sum >= 0 {
				cnt += abs(-1 - a[i] - sum)
				sum = -1
			} else {
				sum += a[i]
			}
		}
	}
	return cnt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
