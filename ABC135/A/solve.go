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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	s := make([]int, 2)
	s[0], s[1] = nextInt(), nextInt()
	sort.Ints(s)

	for i := s[0] + 1; i < s[1]; i++ {
		if abs(i-s[0]) == abs(i-s[1]) {
			fmt.Println(i)
			return
		}
	}

	fmt.Println("IMPOSSIBLE")
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
