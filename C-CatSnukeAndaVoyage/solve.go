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

type Key struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	a := make([]int, m)
	b := make([]int, m)
	ab := make(map[Key]bool)

	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
		ab[Key{a[i], b[i]}] = true
	}

	for i := 1; i <= n; i++ {
		if ab[Key{1, i}] && ab[Key{i, n}] {
			fmt.Println("POSSIBLE")
			return
		}
	}

	fmt.Println("IMPOSSIBLE")

	// land := make([][]int, n+1)
	// for i := 0; i < n+1; i++ {
	// 	land[i] = make([]int, n+1)
	// }
	// // land[0] = 1
	// for i := 0; i < m; i++ {

	// 	land[a[i]][b[i]] = 1
	// }

	// for i := 0; i < m; i++ {
	// 	cost := 0
	// 	if a[i] == 1 {
	// 		prev := a[i]
	// 		next := b[i]
	// 		for {
	// 			cost += land[prev][next]
	// 			prev = next
	// 			next
	// 		}
	// 	}

	// }

	// fmt.Println(land)

	// next := 0
	// cnt := 0
	// can := false
	// for {
	// 	next = land[next]
	// 	cnt++
	// 	// fmt.Printf("next=%d cnt=%d\n", next, cnt)

	// 	if cnt == 3 {
	// 		if next == n {
	// 			can = true
	// 			break
	// 		}
	// 	}

	// 	if cnt > 3 {
	// 		break
	// 	}
	// }

	// if can {
	// 	fmt.Println("POSSIBLE")
	// } else {
	// 	fmt.Println("IMPOSSIBLE")
	// }

}
