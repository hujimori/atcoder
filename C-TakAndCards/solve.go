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

var fact []int

func initFactorial(N int) {
	fact = make([]int, N+1)
	fact[0] = 1
	fact[1] = 1

	for i := N; i > 0; i-- {
		f := 1
		for j := i; j > 0; j-- {
			f *= j
		}
		fact[i] = f
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N, A := nextInt(), nextInt()
	x := make([]int, N)
	for i := 0; i < N; i++ {
		x[i] = nextInt()
	}

	var dp [51][51][3010]int
	dp[0][0][0] = 1
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < 2500; k++ {
				if dp[i][j][k] > 0 {
					dp[i+1][j][k] += dp[i][j][k]
					dp[i+1][j+1][k+x[i]] += dp[i][j][k]
				}
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += dp[N][i][i*A]
	}
	fmt.Println(ans)

	// ans := 0
	// for bit := uint(0); bit < (1 << uint(N)); bit++ {
	// 	slice := []int{}
	// 	for i := uint(0); i < uint(N); i++ {
	// 		if (bit & (1 << i)) == 1<<i {
	// 			slice = append(slice, x[i])
	// 		}
	// 	}
	// 	sum := 0
	// 	for _, ss := range slice {
	// 		sum += ss
	// 	}

	// 	if (len(slice) != 0) && (float64(sum)/float64(len(slice)) == float64(A)) {
	// 		ans++
	// 	}
	// }
	// fmt.Println(ans)
}
