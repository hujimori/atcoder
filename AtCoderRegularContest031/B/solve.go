package main

import "fmt"

var field []string

var landCount int

func main() {

	field = make([]string, 10)

	for i := 0; i < 10; i++ {
		fmt.Scan(&field[i])
	}

	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {

	// 		fmt.Println(string(field[i][j]))
	// 	}
	// }
	// L:

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// fmt.Printf("i = %d, j = %d:", i, j)
			if isOk(i, j) {
				fmt.Println("YES")
				return
			}
		}
	}

	fmt.Println("NO")

}

func isOk(x, y int) bool {
	reached := make([][]bool, 10)
	for i := 0; i < 10; i++ {
		reached[i] = make([]bool, 10)
	}

	// if string(field[x][y]) == "x" {
	dfs(x, y, reached)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if string(field[i][j]) == "o" && !reached[i][j] {
				return false
			}
		}

	}
	// }

	return true
}

var dx = []int{1, 0, 0, -1}
var dy = []int{0, 1, -1, 0}

func dfs(x, y int, reached [][]bool) {

	reached[x][y] = true

	for i := 0; i < len(dx); i++ {

		nx := x + dx[i]
		ny := y + dy[i]

		if nx >= 0 && 10 > nx && ny >= 0 && 10 > ny && !reached[nx][ny] && string(field[nx][ny]) != "x" {
			dfs(nx, ny, reached)
		}

	}

	return
}
