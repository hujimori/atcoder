package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var s = bufio.NewScanner(os.Stdin)

func nextInt() int {
	s.Scan()
	i, e := strconv.Atoi(s.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func solve(H, W int) int {
	var (
		sa int
		sb int
		sc int
	)

	ans := 1000000000000
	for h := 1; h <= H-1; h++ {
		sa = h * W

		for i := 0; i < 2; i++ {
			if i%2 == 0 {
				w := W / 2
				sb = (H - h) * w
				sc = (H - h) * (W - w)
			} else {
				hh := (H - h) / 2
				sb = hh * W
				sc = (H - h - hh) * W
			}

			smax := max(max(sa, sb), sc)
			smin := min(min(sa, sb), sc)

			ans = min(ans, smax-smin)
		}
	}
	return ans
}

func main() {
	s.Split(bufio.ScanWords)

	H, W := nextInt(), nextInt()

	fmt.Println(min(solve(H, W), solve(W, H)))
	// var (
	// 	h int
	// 	w int
	// )
	// if H < W {
	// 	h = H
	// 	w = W
	// } else {
	// 	h = W
	// 	w = H
	// }

	// smax := w * (h / 2)

	// if (h - h/2) < w/2 {

	// }

	// smin := (h - h/2) * (w / 2)

	// fmt.Printf("smax=%d, smin=%d\n", smax, smin)

}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
