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

var slice []int

func main() {
	sc.Split(bufio.ScanWords)

	slice = make([]int, 4)

	for i, n := range strings.Split(nextLine(), "") {
		slice[i], _ = strconv.Atoi(n)
	}

	// fmt.Println(slice)

	if ok, ans := dfs(0, 0, ""); ok {
		fmt.Println(ans[1:len(ans)])
	}

}

func dfs(i, sum int, ex string) (bool, string) {

	if i == 4 {
		return sum == 7, "=7"
	}

	// "+"を使う場合
	if ok, s := dfs(i+1, sum+slice[i], "+"); ok {
		return true, "+" + strconv.Itoa(slice[i]) + s
	}

	// "-"を使う場合
	if ok, s := dfs(i+1, sum-slice[i], "-"); ok {
		return true, "-" + strconv.Itoa(slice[i]) + s
	}

	return false, ""
}
