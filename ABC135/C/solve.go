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

	N := nextInt()
	A := make([]int, N+1)
	B := make([]int, N)

	sum := 0
	for i := 0; i < N+1; i++ {
		A[i] = nextInt()
		sum += A[i]
	}

	for i := 0; i < N; i++ {
		B[i] = nextInt()
	}

	// sum := make([]int, N+2)
	// sum[0] = 0

	// for i := 0; i < N+1; i++ {
	// 	sum[i+1] = sum[i] + A[i]
	// }

	// ans := sum[len(sum)-1]

	// fmt.Println(sum)

	for i := 0; i < N; i++ {
		if A[i]-B[i] < 0 {
			B[i] = B[i] - A[i]
			A[i] = 0

			if A[i+1]-B[i] < 0 {
				B[i] = B[i] - A[i+1]
				A[i+1] = 0
			} else {
				A[i+1] -= B[i]
				B[i] = 0
			}

		} else {
			A[i] -= B[i]
			B[i] = 0

		}
	}
	sum2 := 0
	for i := 0; i < N+1; i++ {
		// fmt.Println(A[i])
		sum2 += A[i]
	}
	// fmt.Println(A)
	fmt.Println(sum - sum2)
}
