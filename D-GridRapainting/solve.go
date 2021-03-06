package main

import (
	"fmt"
)

type P struct {
	X int
	Y int
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
var visted [][]bool

// 各点までの最短距離のスライス
var d [][]int

func main() {
	fmt.Scan(&H, &W)
	sx, sy = 0, 0
	gx, gy = H-1, W-1
	field = make([]string, H)
	visted = make([][]bool, H)
	d = make([][]int, H)

	for i := 0; i < H; i++ {
		// 地形データ初期化
		fmt.Scan(&field[i])

		// 訪問状態を初期化
		visted[i] = make([]bool, W)

		// 各点までの最短距離のスライス初期化
		d[i] = make([]int, W)
	}

	black := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if string(field[h][w]) == "#" {
				black++
			}
		}
	}

	res := bfs()
	if visted[gx][gy] {
		ans := H*W - (res + 1 + black)
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}

}

func bfs() int {
	// 初期位置をキューにpush
	q := Queue{}
	q.push(&P{X: sx, Y: sy})
	// 初期位置のコストを0
	d[sx][sy] = 0
	visted[sx][sy] = true

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	for q.size() > 0 {
		// キューの先頭を取り出す
		p := q.front()
		// fmt.Println(p)
		q.pop()

		// 取り出してきた状態がゴールなら探索を終了
		if p.X == gx && p.Y == gy {
			break
		}

		// 移動4方向をループ
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]

			if 0 <= nx && nx < H && 0 <= ny && ny < W && string(field[nx][ny]) != "#" && !visted[nx][ny] {
				q.push(&P{X: nx, Y: ny})
				d[nx][ny] = d[p.X][p.Y] + 1
				visted[nx][ny] = true

			}
		}

	}

	return d[gx][gy]
}
