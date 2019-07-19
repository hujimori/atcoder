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

var inputs [5]int

func main() {
	sc.Split(bufio.ScanWords)

	for i := 0; i < 5; i++ {
		inputs[i] = nextInt()
	}
	fmt.Println(max(inputs[4]+inputs[3]+inputs[0], inputs[4]+inputs[2]+inputs[1]))

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
