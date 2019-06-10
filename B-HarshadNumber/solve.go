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

	n := nextInt()
	digitSum := 0
	num := n
	for {
		digitSum += num % 10
		num = num / 10

		if num == 0 {
			break
		}
	}

	if n%digitSum == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
