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
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	b := make([]int, n)

	for i := n; i >= 1; i-- {
		// iの倍数の箱に入ってるボールの数の合計
		sum := 0
		for j := i + i; j <= n; j += i {
			sum ^= b[j-1]
		}
		b[i-1] = sum ^ a[i-1]

	}

	ans := make([]int, 0, len(a))
	for i, e := range b {
		if e == 1 {
			ans = append(ans, i+1)
		}
	}

	fmt.Println(len(ans))
	if len(ans) != 0 {
		for _, a := range ans {
			fmt.Printf("%d ", a)
		}
		fmt.Println()

	}

}

// func main() {
// 	sc.Split(bufio.ScanWords)

// 	n := nextInt()
// 	a := make([]int, n)
// 	b := make([]int, n+1)
// 	for i := 0; i < n; i++ {
// 		a[i] = nextInt()
// 	}

// 	m := map[int]int{}
// 	for i := 1; i <= n; i++ {
// 		m[i] = n / i
// 	}

// 	for i := n; i >= 1; i-- {
// 		if m[i] == 1 {
// 			b[i] = 1
// 			continue
// 		} else {
// 			b[i] = 0
// 			continue
// 		}

// 	}

// }
