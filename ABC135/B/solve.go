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

	N := nextInt()
	p := make([]int, N)
	pp := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = nextInt()
		pp[i] = p[i]
	}

	sort.Ints(pp)

	cnt := 0
	for i := 0; i < N; i++ {
		if p[i] != pp[i] {
			cnt++
		}
	}

	if cnt <= 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
