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

	S := strings.Split(nextLine(), "")

	var smax string
	for _, ss := range S {
		if ss == "?" {
			smax += "9"
		} else {
			smax += ss
		}
	}

	var smin string
	flag := false
	for _, ss := range S {
		if ss != "?" {
			flag = true
		}

		if !flag {
			continue
		}

		if ss == "?" {
			smin += "0"
		} else {
			smin += ss
		}
	}

	max, _ := strconv.Atoi(smax)
	min, _ := strconv.Atoi(smin)

	fmt.Println(max)
	fmt.Println(min)

	table := make([]string, 100010)
	for i := 0; i < 100010; i++ {
		table[i] = strconv.Itoa(i)
	}

	// sum := max - min + 1
	// ans :=

	ans := 0
	mod := 1000000007
	for i := min; i <= max; i++ {
		if i%13 == 5 {
			ans++
		}
	}

	fmt.Println(ans % mod)
}
