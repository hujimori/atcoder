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
	N, M := nextInt(), nextInt()

	s := make([][]int, M)
	for i := 0; i < M; i++ {
		k := nextInt()
		for j := 0; j < k; j++ {
			a := nextInt() - 1
			s[i] = append(s[i], a)
		}
	}

	p := make([]int, M)
	for i := 0; i < M; i++ {
		p[i] = nextInt()
	}

	res := 0
	for bit := uint(0); bit < (1 << uint(N)); bit++ {
		ok := true
		for i := 0; i < M; i++ {
			con := 0
			for _, v := range s[i] {
				if (bit & (1 << uint(v))) == (1 << uint(v)) {
					con++
				}
			}
			if con%2 != p[i] {
				ok = false
			}
		}
		if ok {
			res++
		}
	}

	fmt.Println(res)
}
