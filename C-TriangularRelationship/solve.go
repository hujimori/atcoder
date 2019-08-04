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

	n, k := nextInt(), nextInt()
	num := make([]int, k)
	for i := 1; i <= n; i++ {
		num[i%k]++
	}

	res := 0
	for a := 0; a < k; a++ {
		b := (k - a) % k
		c := (k - a) % k
		if (b+c)%k != 0 {
			continue
		}
		res += num[a] * num[b] * num[c]
	}

	fmt.Println(res)

}
