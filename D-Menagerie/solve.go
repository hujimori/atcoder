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
	var N int
	var S string

	fmt.Scan(&N)
	fmt.Scan(&S)

	s := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == 'o' {
			s[i] = 0
		} else {
			s[i] = 1
		}
	}

	for t0 := 0; t0 < 2; t0++ {
		for t1 := 0; t1 < 2; t1++ {
			t := make([]int, N)
			t[0] = t0
			t[1] = t1

			for i := 2; i < N; i++ {
				t[i] = (t[i-2] ^ t[i-1] ^ s[i-1])
			}
			if (s[N-1] == (t[N-2] ^ t[N-1] ^ t[0])) && (s[0] == (t[N-1] ^ t[0] ^ t[1])) {
				ans := ""
				for _, tt := range t {
					if tt == 0 {
						ans += "S"
					} else {
						ans += "W"
					}
				}
				fmt.Println(ans)
				return
			}

		}
	}

	fmt.Println(-1)

}
