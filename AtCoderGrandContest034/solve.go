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
	n := nextInt()
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()

	s := strings.Split(nextLine(), "")

	s[a] = "A"
	s[b] = "B"

	fmt.Println("Input")
	fmt.Printf("N:%d,A:%d,B:%d,C:%d,D:%d\n", n, a, b, c, d)
	fmt.Println(s)

	for s[c] == "A" && s[d] == "B" {
		// Aが動けるか調べる
		if a+1 < n && (s[a+1] == "." ) 
	}
}
