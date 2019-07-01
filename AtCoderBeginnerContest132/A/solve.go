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

func main() {
	sc.Split(bufio.ScanWords)

	var s string
	fmt.Scan(&s)

	slice := strings.Split(s, "")
	m := map[string]int{}

	for _, ss := range slice {
		m[ss]++
	}

	if len(m) == 2 {
		for _, v := range m {
			if v != 2 {
				fmt.Println("No")
				return
			}
		}
		fmt.Println("Yes")
		return
	}

	fmt.Println("No")

}
