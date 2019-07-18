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
	r, c, k := nextInt(), nextInt(), nextInt()
	n := nextInt()

	rs := make([]int, n)
	cs := make([]int, n)
	rcnt := make([]int, r)
	ccnt := make([]int, c)
	for i := 0; i < n; i++ {
		rs[i] = nextInt() - 1
		cs[i] = nextInt() - 1
		rcnt[rs[i]]++
		ccnt[cs[i]]++
	}

	m := map[int]int{}
	for i := 0; i < c; i++ {
		m[ccnt[i]]++
	}

	ans := 0
	for i := 0; i < r; i++ {
		ans += m[k-rcnt[i]]
	}

	for i := 0; i < n; i++ {
		d := rcnt[rs[i]] + ccnt[cs[i]]

		if d == k {
			ans--
		}
		if d == k+1 {
			ans++
		}
	}

	fmt.Println(ans)

	// for i := 0; i < r; i++ {
	// 	for j := 0; j < c; j++ {
	// 		if field[i][j] {
	// 			line[i]++
	// 		}
	// 	}
	// }

	// for i := 0; i < c; i++ {
	// 	for j := 0; j < r; j++ {
	// 		if field[j][i] {
	// 			row[i]++
	// 		}
	// 	}
	// }

	// ans := 0
	// for i := 0; i < r; i++ {
	// 	for j := 0; j < c; j++ {
	// 		if field[i][j] {
	// 			if (line[i] + row[j] - 1) == k {
	// 				ans++
	// 			}
	// 		} else {
	// 			if (line[i] + row[j]) == k {
	// 				ans++
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(ans)

}
