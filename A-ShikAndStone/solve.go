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

	H, W := nextInt(), nextInt()
	field := make([]string, H)

	for i := 0; i < H; i++ {
		field[i] = nextLine()
	}

	cnt := 0
	for j := 0; j < H; j++ {
		for i := 0; i < W; i++ {
			if field[j][i] == '#' {
				cnt++
			}
		}
	}
	if H+W-1 == cnt {
		fmt.Println("Possible")
	} else {
		fmt.Println("Impossible")
	}

}
