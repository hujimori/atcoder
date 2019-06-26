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

var v []int
var n int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	v = make([]int, n)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		v[i] = nextInt()
		m[v[i]]++
	}

	if len(m) == 1 {
		fmt.Println(n / 2)
		return
	}

	oddMax := -1
	evenMax := -1

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if evenMax < m[v[i]] {
				evenMax = v[i]
			}
		} else {
			if oddMax < m[v[i]] {
				oddMax = v[i]
			}
		}
	}

	// fmt.Printf("evenmax=%d, oddmax=%d\n", evenMax, oddMax)

	cnt := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			if v[i] != evenMax {
				cnt++
			}
		} else {
			if v[i] != oddMax {
				cnt++
			}
		}
	}
	fmt.Println(cnt)

	// fmt.Println(m)

	// cnt := 0
	// for i := 0; i < len(v)/2; i++ {
	// 	cnt += dfs(i, 0, v)
	// }
	//  else {
	// 		fmt.Println(cnt)
	// 	}
	// fmt.Println(cnt)
}

func dfs(i, cnt int, v []int) int {
	if i >= n {
		return cnt
	}

	var res int

	if i+2 < n && v[i] != v[i+2] {
		tmp1 := v[i]
		tmp2 := v[i+2]

		v[i] = tmp2
		chmin(&res, dfs(i+2, cnt+1, v))

		v[i] = tmp1
		v[i+2] = tmp1
		chmin(&res, dfs(i+2, cnt+1, v))
	} else {
		chmin(&res, dfs(i+2, cnt, v))
	}

	fmt.Println(res)
	return res
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}
