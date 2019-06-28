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

var n int
var k int
var a []int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	k = 2
	a = make([]int, n)
	sum := 1
	even := 0
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if a[i]%2 == 0 {
			even++
		}
		sum *= a[i]
	}
	fmt.Println()

	cnt := 0
	for sum != 0 {
		sum /= 2
		cnt++
		if cnt >= n && even >= n/2 {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")

	// fmt.Println(e)

	// var slice [][]int

	// for bit := uint(0); bit < (1 << uint(n)); bit++ {
	// 	s := []int{}
	// 	for i := uint(0); i < uint(n); i++ {
	// 		if bit&(1<<i) == 1<<i {
	// 			s = append(s, a[i])
	// 		}
	// 	}
	// 	slice = append(slice, s)
	// }

	// solve()
}

func nextCombination(sub int) int {
	x := sub & -sub
	y := sub + x
	return (((sub & ^y) / x) >> 1) | y
}

func solve() {

	bit := (1<<uint(k) - 1)
	slice := [][]int{}
	for ; bit < (1 << uint(n)); bit = nextCombination(bit) {
		s := []int{}
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) == 1<<uint(i) {
				s = append(s, a[i])
			}
		}
		slice = append(slice, s)
	}
	cnt := 0
	for _, s := range slice {
		if s[0]*s[1]%4 == 0 {
			fmt.Println(s)
			cnt++
		}
	}

	if cnt >= n-1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
