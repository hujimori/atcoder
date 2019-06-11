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

	str := nextLine()
	m := map[string]int{}
	// fmt.Println(strings.Split(s, ""))
	// sort.Sort(sort.StringSlice(strings.Split(s, "")))
	// fmt.Println(s)

	for _, key := range str {
		if _, ok := m[string(key)]; !ok {
			m[string(key)] = 1
		} else {
			m[string(key)]++
		}
	}

	fmt.Println(m)
}
