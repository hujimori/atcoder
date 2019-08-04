package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

const INF = 1000000001000

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = nextLine()
	}

	alpha := map[string]int{}

	for _, abc := range "abcdefghijklmnopqrstuvwxyz" {
		alpha[string(abc)] = INF
	}

	for _, ss := range s {
		m := map[string]int{}
		for _, abc := range "abcdefghijklmnopqrstuvwxyz" {
			m[string(abc)] = 0
		}

		for _, sss := range ss {
			m[string(sss)]++
		}

		for k, v := range m {
			alpha[k] = min(v, alpha[k])
		}
	}
	var ans []string
	for k, v := range alpha {
		if v != 0 {
			ans = append(ans, strings.Repeat(k, v))
		}
	}
	sort.Strings(ans)
	fmt.Println(strings.Join(ans, ""))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
