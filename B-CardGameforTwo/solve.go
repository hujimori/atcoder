package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N := nextInt()
	A := make([]int, N)

	for i := 0; i < N; i++ {
		A = append(A, nextInt())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	Alice := 0
	Bob := 0
	n := 0
	for len(A) > 0 {

		if len(A) > 0 {
			n = A[0]
			A = A[1:]
			Alice += n
		}

		if len(A) > 0 {
			n = A[0]
			A = A[1:]
			Bob += n
		}
	}

	fmt.Println(Alice - Bob)

}
