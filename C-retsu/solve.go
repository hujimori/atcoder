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
	n, k := nextInt(), nextInt()
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = nextInt()
	}

	for i := 0; i < n; i++ {
		if s[i] == 0 {
			fmt.Println(n)
			return
		}
	}

	// mul := make([]int, n+1)
	// mul[0] = 1
	// for i := 0; i < n; i++ {
	// 	mul[i+1] = mul[i] * s[i]
	// }
	// fmt.Println(mul)

	res := 0
	right := 0
	mul := 1
	for left := 0; left < n; left++ {

		for right < n && mul*s[right] <= k {
			mul *= s[right]
			right++
		}

		res = max(res, right-left)
		if right == left {
			right++
		} else {
			mul /= s[left]
		}
	}

	fmt.Println(res)
	// ans := -1
	// mul := 1

	// for left := 0; left < n; left++ {
	// 	right := left

	// 	for right < n && mul*s[right] <= k {
	// 		mul *= s[right]
	// 		right++

	// 	}

	// 	ans = max(ans, right-left)
	// }

	// for {
	// 	if sum*s[right] <= k {
	// 		sum *= s[right]
	// 		right++
	// 	} else {
	// 		ans = max(ans, right-left)
	// 		sum /= s[left]
	// 		left++
	// 	}

	// 	if right == len(s) && left == len(s) {
	// 		break
	// 	}

	// }

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
