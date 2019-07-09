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

var (
	n int
	c []int
	s []int
	f []int
)

func main() {
	sc.Split(bufio.ScanWords)

	n = nextInt()
	c = make([]int, n)
	s = make([]int, n)
	f = make([]int, n)

	for i := 0; i < n-1; i++ {
		c[i], s[i], f[i] = nextInt(), nextInt(), nextInt()
	}

	for i := 0; i < n-1; i++ {
		t := s[i] + c[i]
		for j := i + 1; j < n; j++ {
			if t <= s[j] {
				t += s[j] - t
			}

			if t > s[j] {
				if f[j] != 0 && (t-s[j])%f[j] != 0 {
					t += (t/f[j]+1)*f[j] - t
				}

				// if f[j] != 0 && t%f[j] != 0 {

				// }
			}

			t += c[j]
		}
		fmt.Println(t)
	}
	fmt.Println(0)
}
