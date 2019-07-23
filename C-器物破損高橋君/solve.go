package main

import (
	"fmt"
)

type P struct {
	X     int
	Y     int
	Break int
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
	sy int
	sx int
	gy int
	gx int
)

// 地形データ
var field []string

// 訪問状態
var visted [][][]bool

func main() {
	fmt.Scan(&H, &W)

	field = make([]string, H)
	visted = make([][][]bool, H)

	for i := 0; i < H; i++ {
		// 地形データ初期化
		fmt.Scan(&field[i])

		// 訪問状態を初期化
		visted[i] = make([][]bool, W)
		for j := 0; j < W; j++ {
			visted[i][j] = make([]bool, 3)
		}
	}

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if string(field[h][w]) == "s" {
				sx = h
				sy = w
			}

			if string(field[h][w]) == "g" {
				gx = h
				gy = w
			}
		}
	}

	if bfs() {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

var INF = 1000000000000

func bfs() bool {
	// 初期位置をキューにpush
	q := Queue{}
	q.push(&P{X: sx, Y: sy, Break: 0})
	visted[sx][sy][0] = true

	dx := []int{1, 0, -1, 0} // X軸
	dy := []int{0, 1, 0, -1} // Y軸

	for q.size() > 0 {
		// キューの先頭を取り出す
		p := q.front()
		q.pop()

		// 取り出してきた状態がゴールなら探索を終了
		if p.X == gx && p.Y == gy {
			break
		}

		// 移動4方向をループ
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]
			nb := p.Break
			if 0 <= nx && nx < H && 0 <= ny && ny < W {
				if string(field[nx][ny]) == "#" && nb == 2 {
					continue
				}
				if visted[nx][ny][nb] {
					continue
				}

				if string(field[nx][ny]) == "#" {
					q.push(&P{X: nx, Y: ny, Break: nb + 1})
					visted[nx][ny][nb+1] = true
				} else {
					q.push(&P{X: nx, Y: ny, Break: nb})
					visted[nx][ny][nb] = true
				}
			}
		}

	}

	for i := 0; i < 3; i++ {
		if visted[gx][gy][i] {
			return true
		}
	}

	return false
}
