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
	n, k := nextInt(), nextInt()
	for i := 1; i <= k; i++ {
		fmt.Println(perm(i, i) * perm(n-i, n-i))
		// fmt.Println(solve(i, i) * solve(n-k+1, k+1))
	}
	// fmt.Println(combi(5, 3))

}

func combi(n, r int) int {
	if r == 0 {
		return 1
	}

	return (n - r + 1) * combi(n, r-1) / r
}

func perm(n, r int) int {
	p := 1
	for i := 0; i < r; i++ {
		p *= (n - 1)
	}
	return p
}
func nextCombination(sub int) int {
	x := sub & -sub
	y := sub + x
	return (((sub & ^y) / x) >> 1) | y
}

func solve(n, k int) int {
	fmt.Printf("%d, %d\n", n, k)
	// s := [][]int{}
	cnt := 0
	bit := (1<<uint(k) - 1)
	for ; bit < (1 << uint(n)); bit = nextCombination(bit) {

		slice := []int{}
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) == 1<<uint(i) {
				slice = append(slice, i)

				cnt++

			}
		}
		// fmt.Printf("%d: {", bit)
		// for i := 0; i < len(slice); i++ {
		// 	fmt.Printf("%d ", slice[i])
		// }
		// fmt.Print("}\n")

	}
	return cnt
}
