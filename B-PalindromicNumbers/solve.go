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

	a := nextInt()
	b := nextInt()
	count := 0
	for i := a; i <= b; i++ {
		str := strconv.Itoa(i)
		if str[0] == str[len(str)-1] && str[1] == str[len(str)-2] {
			count++
			// fmt.Println(str)
		}
	}

	fmt.Println(count)
}
