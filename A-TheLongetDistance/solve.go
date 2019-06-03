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

type Point struct {
	x float64
	y float64
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2.0) + math.Pow(p1.y-p2.y, 2.0))
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	slice := make([]Point, n)
	for i := 0; i < n; i++ {
		slice[i] = Point{
			x: float64(nextInt()),
			y: float64(nextInt()),
		}
	}
	fmt.Println(slice)

	max := math.Inf(-1)
	res := 0.0
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			res = distance(slice[i], slice[j])
			if max < res {
				max = res
			}
		}
	}
	fmt.Println(max)
}
