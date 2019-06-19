package main

import (
	"fmt"
)

var field []string
var reached [][]bool

var h int
var w int

func main() {

	fmt.Scan(&h, &w)

	field = make([]string, h)
	reached = make([][]bool, h)

	for i := 0; i < h; i++ {
		fmt.Scan(&field[i])
		reached[i] = make([]bool, w)

	}

	var startx, starty, goalx, goaly int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {

			if field[i][j] == 's' {
				startx = i
				starty = j
			}

			if field[i][j] == 'g' {
				goalx = i
				goaly = j
			}
		}
	}

	dfs(startx, starty)

	if reached[goalx][goaly] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(x, y int) {
	if x < 0 || w <= x || y < 0 || h <= y || field[x][y] == '#' {
		return
	}

	if reached[x][y] {
		return
	}

	reached[x][y] = true

	dfs(x+1, y)
	dfs(x-1, y)
	dfs(x, y+1)
	dfs(x, y-1)
}
