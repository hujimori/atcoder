package main

import (
	"bufio"
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

type Beaker struct {
	water int
	sugar int
	cap   int
}

func (b *Beaker) isOver() bool {
	if b.water+b.sugar > b.cap {
		return true
	}

	return false
}

func (b *Beaker) addWater(w int) {
	b.water += w
}

func (b *Beaker) addSugar(s int) {
	b.sugar += s
}

func main() {
	sc.Split(bufio.ScanWords)

	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()
	e := nextInt()
	f := nextInt()

	// operation=0:ビーカーに対して何もしない
	// operation=1:ビーカーに水を100A[g]入れる
	// operation=2:ビーカーに水を100B[g]入れる
	// operation=3:ビーカーに砂糖をc[g]入れる
	// operation
	opration := []int{0, 1, 2, 3, 4}

	for _, op := range opration {

	}

}
