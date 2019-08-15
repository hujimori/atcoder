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

var dp [100][2]int

const MAX_DIGIT = 50

func main() {
	sc.Split(bufio.ScanWords)

	N, K := nextInt(), nextInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 2; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = 0
	for d := 0; d < MAX_DIGIT; d++ {
		mask := 1 << uint((MAX_DIGIT - d - 1))

		num := 0
		for i := 0; i < N; i++ {
			if (A[i] & mask) == mask {
				num++
			}
		}

		cost0 := mask * num
		cost1 := mask * (N - num)

		if dp[d][1] != -1 {
			chmax(&dp[d+1][1], dp[d][1]+max(cost0, cost1))
		}

		if dp[d][0] != -1 {
			if (K & mask) == mask {
				chmax(&dp[d+1][1], dp[d][0]+cost0)
			}
		}

		if dp[d][0] != -1 {
			if (K & mask) == mask {
				chmax(&dp[d+1][0], dp[d][0]+cost1)
			} else {
				chmax(&dp[d+1][0], dp[d][0]+cost0)
			}
		}
	}

	fmt.Println(max(dp[MAX_DIGIT][0], dp[MAX_DIGIT][1]))

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
