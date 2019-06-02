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

var grid [][]string

func main() {
	sc.Split(bufio.ScanWords)

	H := nextInt()
	W := nextInt()

	grid = make([][]string, H+2)

	for i := 0; i < H+2; i++ {
		grid[i] = make([]string, W+2)
		for j := 0; j < W+2; j++ {
			grid[i][j] = "@"
		}
	}

	for i := 1; i < H+1; i++ {
		line := nextLine()
		for j, l := range strings.Split(line, "") {
			grid[i][j+1] = l
		}
	}

	// fmt.Println(grid)

	// for i := 0; i < H+2; i++ {
	// 	for j := 0; j < W+2; j++ {
	// 		fmt.Print(grid[i][j])
	// 	}
	// 	fmt.Println()
	// }

	can := true

	for i := 1; i < H+1; i++ {
		for j := 1; j < W+1; j++ {
			if grid[i][j] == "." {
				continue
			}

			if !(grid[i+1][j] == "#" || grid[i-1][j] == "#" || grid[i][j-1] == "#" || grid[i][j+1] == "#") {
				can = false
			}

		}
	}

	if can {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
