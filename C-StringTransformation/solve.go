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

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")

	ms := map[string]string{}
	mt := map[string]string{}
	ok := true
	for i := 0; i < len(ss); i++ {

		if len(ms[ss[i]]) != 0 {
			if ms[ss[i]] != tt[i] {
				ok = false
			}
		}

		if len(mt[tt[i]]) != 0 {
			if mt[tt[i]] != ss[i] {
				ok = false
			}
		}

		ms[ss[i]] = tt[i]
		mt[tt[i]] = ss[i]
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
