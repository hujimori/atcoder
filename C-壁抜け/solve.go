package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type P struct {
	X     int
	Y     int
	Cost  int
	Black int
}

type Queue struct {
	p []*P
}

func (q *Queue) push(p *P) {
	q.p = append(q.p, p)
}

func (q *Queue) pop() {
	q.p = q.p[1:]
}

func (q *Queue) front() *P {
	return q.p[0]
}

func (q *Queue) size() int {
	return len(q.p)
}

var (
	H  int
	W  int
	T  int
	sy int
	sx int
	gy int
	gx int
)

// 地形データ
var field []string

// 訪問状態
// var

// 各点までの最短距離のスライス
// var d [][]int

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	// sc.Split(bufio.ScanWords)
	fmt.Scan(&H, &W, &T)

	field = make([]string, H)
	// visted = make([][]bool, H)
	// d = make([][]int, H)

	for i := 0; i < H; i++ {
		// 地形データ初期化
		fmt.Scan(&field[i])

		// 訪問状態を初期化
		// visted[i] = make([]bool, W)

		// 各点までの最短距離の初期化
		// d[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if string(field[i][j]) == "S" {
				sx = i
				sy = j
			}
			if string(field[i][j]) == "G" {
				gx = i
				gy = j
			}
		}
	}
	solve()
	// res := bfs(8)
	// if res <= T {
	// 	fmt.Println(res)
	// 	return
	// }
}

func solve() {
	lb := 0
	ub := T

	for ub-lb > 1 {
		mid := (lb + ub) / 2
		if bfs(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	fmt.Println(lb)
}

// func C(x int) bool {
// 	bfs(x)
// 	// fmt.Printf("x=%d, res=%d\n", x, res)
// 	for i := 0; i < 110; i++ {
// 		if
// 	}

// 	if res <= T {
// 		return true
// 	}
// 	return false
// }

var INF = 100000000000

func bfs(x int) bool {
	// 初期位置をキューにpush
	q := Queue{}
	q.push(&P{X: sx, Y: sy, Black: 0, Cost: 0})

	visted := [20][20][220]int{}
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			for k := 0; k < 220; k++ {
				visted[i][j][k] = INF
			}
		}
	}
	visted[sx][sy][0] = 0

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for q.size() > 0 {
		// キューの先頭を取り出す
		p := q.front()
		q.pop()

		// 移動4方向をループ
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]
			var ncost int
			var nblack int
			if 0 <= nx && nx < H && 0 <= ny && ny < W {

				if string(field[nx][ny]) == "." || string(field[nx][ny]) == "S" || string(field[nx][ny]) == "G" {
					ncost = p.Cost + 1
					nblack = p.Black + 0
				} else {
					ncost = p.Cost + x
					nblack = p.Black + 1
				}

				if nblack >= 220 {
					continue
				}

				if visted[nx][ny][nblack] != INF {
					continue
				}
				visted[nx][ny][nblack] = ncost

				if nx == gx && ny == gy {
					continue
				}
				q.push(&P{X: nx, Y: ny, Cost: ncost, Black: nblack})
			}
		}

	}

	for i := 0; i < 110; i++ {
		if visted[gx][gy][i] <= T {
			return true
		}
	}
	return false

	// return d[gx][gy]
}
