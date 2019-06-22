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
	sc.Split(bufio.ScanWords)

	a := nextInt()
	b := nextInt()
	x := nextInt()

	var answer int64

	answer = int64(b/x) + 1

	if a == 0 {
		answer += 0
	} else {
		answer -= int64((a-1)/x) + 1
	}

	// for i := a; i <= b; i++ {
	// 	if i%x == 0 {
	// 		answer++
	// 	}
	// }

	fmt.Println(answer)
}
