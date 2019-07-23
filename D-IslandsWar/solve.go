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

type AB struct {
	A int
	B int
}

type ABS []AB

func (abs ABS) Len() int {
	return len(abs)
}

func (abs ABS) Swap(i, j int) {
	abs[i], abs[j] = abs[j], abs[i]
}

func (abs ABS) Less(i, j int) bool {
	return abs[i].B < abs[j].B
}

func main() {
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	slice := make([]AB, m)
	for i := 0; i < m; i++ {
		slice[i] = AB{A: nextInt(), B: nextInt()}
	}
	n++

	sort.Sort(ABS(slice))

	res := 0
	end := 0 // 最後に選んだ島の位置
	for i := 0; i < m; i++ {
		if slice[i].A >= end {
			end = slice[i].B
			res++
		}
	}

	fmt.Println(res)
}
