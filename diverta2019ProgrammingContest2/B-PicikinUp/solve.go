package main

import (
	"bufio"
	"fmt"
	"math"
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

type XY struct {
	x int
	y int
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	xy := make([]XY, n)

	for i := 0; i < n; i++ {
		xy[i].x = nextInt()
		xy[i].y = nextInt()
	}

	min := math.Inf(1)

	for i := 0; i < n; i++ {
		cost := 0
		p := 0
		q := 0
		prevx := xy[i].x
		prevy := xy[i].y
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			if p == 0 && q == 0 {
				p = xy[j].x - xy[i].x
				q = xy[j].y - xy[i].y
				if p == 0 && q == 0 {
					break
				}
			}

			fmt.Printf("%d=%d-%d %d=%d-%d\n", p, xy[j].x, xy[i].x, q, xy[j].y, xy[i].y)

			if (xy[j].x-prevx) == p && (xy[j].y-prevy) == q {
				cost += 0
			} else {
				cost++
			}
			prevx = xy[j].x
			prevy = xy[j].y

			// fmt.Printf("cost=%d\n", cost)
		}
		fmt.Printf("最終的なcost=%d\n", cost)

		if cost == 0 {
			cost = 1
		}

		if min > float64(cost) {
			min = float64(cost)
		}
	}

	fmt.Println(min)
}
