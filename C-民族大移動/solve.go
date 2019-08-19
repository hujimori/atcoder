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

	N, D, K := nextInt(), nextInt(), nextInt()
	L := make([]int, D)
	R := make([]int, D)
	for i := 0; i < D; i++ {
		L[i], R[i] = nextInt(), nextInt()
	}

	S := make([]int, K)
	T := make([]int, K)
	for i := 0; i < K; i++ {
		S[i], T[i] = nextInt(), nextInt()
	}

	for k := 0; k < K; k++ {
		now := S[k]

		for d := 0; d < D; d++ {
			if L[d] <= now && now <= R[d] {
				if T[k] <= now {
					now = L[d]
					if now <= T[k] {
						fmt.Println(d + 1)
						break
					}
				} else {
					now = R[d]
					if T[k] <= R[d] {
						fmt.Println(d + 1)
						break
					}
				}
			}
		}
	}

	// for d := 0; d < D; d++ {
	// 	for k := 0; k < K; k++ {

	// 		if S[k] == T[k] {
	// 			fmt.Println(d)

	// 		}

	// 		if L[d] <= S[k] && S[k] <= R[d] {

	// 			if S[k] < T[k] {
	// 				S[k] += min(R[d], T[k]) - S[k]
	// 			} else {
	// 				S[k] -= S[k] - max(L[d], T[k])
	// 			}

	// 		}

	// 	}
	// }
	N++
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
