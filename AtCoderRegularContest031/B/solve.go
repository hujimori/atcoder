package main

import "fmt"

var field []string
var reached [][]bool

var landCount int
var h int
var w int

func main() {

	h := 10
	w := 10

	field = make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&field[i])
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if field[j][j] == 'o' {
				landCount++
			}
		}
	}

	// L:

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if field[i][j] == 'x' {
				fmt.Printf("dfs=%d\n", dfs(i, j, 0))
				// {
				// fmt.Printf("landCount=%d\n", landCount)
				// fmt.Println("YES")

				// break L
				// }
			}
		}
	}

	fmt.Println("No")

}

func dfs(x, y, count int) int {

	if count != 0 && field[x][y] == 'x' {
		fmt.Println(count)

		return count
	}

	if field[x][y] == 'o' {
		count++
	}

	if 0 <= x && x < w && 0 <= y && y < h {
		return max(max(dfs(x+1, y, count), dfs(x-1, y, count)), max(dfs(x, y+1, count), dfs(x, y-1, count)))
	}

	return count
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
