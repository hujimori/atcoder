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

	slice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		slice[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			slice[i][j] = nextInt()
		}
	}

	// slice := [3][3]int{
	// 	{1, 0, 1},
	// 	{2, 1, 2},
	// 	{1, 0, 1},
	// }

	// fmt.Println(slice)

	a := make([]int, 3)
	b := make([]int, 3)
	a[0] = 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[j] = slice[i][j] - a[i]
			a[i+1] = slice[i+1][j] - b[j]
		}
	}

	// fmt.Println(a)
	// fmt.Println(b)

	right := true

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if slice[i][j] != (a[i] + b[j]) {
				right = false
			}
		}
	}

	if right {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
