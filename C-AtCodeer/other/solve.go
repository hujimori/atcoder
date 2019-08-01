// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// var sc = bufio.NewScanner(os.Stdin)

// func nextInt() int {
// 	sc.Scan()
// 	i, e := strconv.Atoi(sc.Text())
// 	if e != nil {
// 		panic(e)
// 	}
// 	return i
// }

// func main() {
// 	sc.Split(bufio.ScanWords)

// 	N := nextInt()
// 	T := make([]int, N)
// 	A := make([]int, N)
// 	for i := 0; i < N; i++ {
// 		T[i], A[i] = nextInt(), nextInt()
// 	}

// 	x := T[0]
// 	y := A[0]

// 	for i := 1; i < N; i++ {
// 		k1 := T[i] - x
// 		k2 := A[i] - y

// 		if k1 == 0 && k2 == 0 {
// 			continue
// 		}

// 		l := 0
// 		r := 100000000000000000
// 		mid := 0
// 		t1 := k1
// 		t2 := k2
// 		for r-l > 1 {
// 			mid = (r + l) / 2
// 			k1 = T[i]*mid - x
// 			k2 = A[i]*mid - y

// 			if k1 == 0 && k2 == 0 {
// 				t1 = 0
// 				t2 = 0
// 				break
// 			}
// 			if k1 >= 0 && k2 >= 0 {
// 				t1 = k1
// 				t2 = k2
// 				r = mid
// 			} else {
// 				l = mid
// 			}
// 		}
// 		x += t1
// 		y += t2

// 	}

// 	fmt.Println(x + y)
// }
