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

func main() {
	sc.Split(bufio.ScanWords)

	N := nextInt()

	depth := 0
	for n := N; n > 0; n /= 2 {
		depth++
	}

	// fmt.Println(depth)
	player := true
	x := 1

	for {

		if depth%2 == 0 {
			// 木の高さが偶数なら
			// Aは左に Bは右に進むのが最善
			if player {
				x = 2 * x
			} else {
				x = 2*x + 1
			}

			if x > N {
				if player {
					fmt.Println("Aoki")

				} else {
					fmt.Println("Takahashi")
				}
				return
			}
			if player {
				player = false
			} else {
				player = true
			}
		} else {
			// 木の高さが奇数なら
			// Aは右に Bは左に進むのが最善
			if player {
				x = 2*x + 1
			} else {
				x = 2 * x
			}

			if x > N {
				if player {
					fmt.Println("Aoki")

				} else {
					fmt.Println("Takahashi")
				}
				return
			}

			if player {
				player = false
			} else {
				player = true
			}
		}

	}

	// player := true

	// for {

	// 	N -= N/2 + N%2

	// 	if N <= 0 {
	// 		if player {
	// 			fmt.Println("Aoki")
	// 			return
	// 		} else {
	// 			fmt.Println("Takahashi")
	// 			return
	// 		}

	// 	}

	// 	player = false

	// }
}
