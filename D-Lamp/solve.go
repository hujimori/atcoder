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

	H, W := nextInt(), nextInt()
	S := make([]string, H)
	// for i := 0; i < H; i++ {
	// 	S[i] = make([]string, W)
	// }

	for i := 0; i < H; i++ {
		S[i] = nextLine()
	}

	var left, right, up, down [2100][2100]int

	for i := 0; i < H; i++ {
		cur := 0
		for j := 0; j < W; j++ {
			if S[i][j] == '#' {
				cur = 0
			} else {
				cur++
			}
			left[i][j] = cur

		}
	}

	for i := 0; i < H; i++ {
		cur := 0
		for j := W - 1; j >= 0; j-- {
			if S[i][j] == '#' {
				cur = 0
			} else {
				cur++
			}
			right[i][j] = cur
		}
	}

	for j := 0; j < W; j++ {
		cur := 0
		for i := 0; i < H; i++ {
			if S[i][j] == '#' {
				cur = 0
			} else {
				cur++
			}
			up[i][j] = cur
		}
	}

	for j := 0; j < W; j++ {
		cur := 0
		for i := H - 1; i >= 0; i-- {

			if S[i][j] == '#' {
				cur = 0
			} else {
				cur++
			}
			down[i][j] = cur
		}
	}

	res := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			res = max(res, left[i][j]+right[i][j]+up[i][j]+down[i][j]-3)
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
