package main

import (
	"bufio"
	"fmt"
	"os"
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

var end int
var slice []string

func main() {
	sc.Split(bufio.ScanWords)

	slice = strings.Split(nextLine(), "")
	end = len(slice)

	ans := dfs(0, "")
	fmt.Println(ans)
}

func dfs(i int, ex string) string {
	if i == end {
		return ex
	}

	// i番目の後ろに+を挿入する場合
	ex += dfs(i+1, slice[i]+"+")
	fmt.Println(ex)

	// i番目の後ろに+を挿入しない場合
	ex += dfs(i+1, slice[i])

	fmt.Println(ex)

	return ex
}
