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

const MOD = 1000000007

var dp [100005][13]int
var n int

func main() {
	var s string

	fmt.Scan(&s)
	n := len(s)

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		var c int
		if s[i] == '?' {
			c = -1
		} else {
			c = int(s[i] - '0')
		}

		for j := 0; j < 10; j++ {
			if c != -1 && c != j {
				continue
			}
			for ki := 0; ki < 13; ki++ {
				dp[i+1][(ki*10+j)%13] += dp[i][ki]
			}
		}

		for j := 0; j < 13; j++ {
			dp[i+1][j] %= MOD
		}
	}
	res := dp[n][5]

	fmt.Println(res)
}
