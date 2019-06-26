package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

type PY struct {
	Id     int
	P      int    // 所属している県
	Y      int    // 誕生都市
	Rank   int    // 誕生時の順番
	Serial string // シリアル
}

func (p *PY) NewSerial() {

	upper := strings.Repeat("0", (6-len(strconv.Itoa(p.P)))) + strconv.Itoa(p.P)
	lower := strings.Repeat("0", (6-len(strconv.Itoa(p.Rank)))) + strconv.Itoa(p.Rank)
	p.Serial = upper + lower
}

type PYS []PY

func (p PYS) Len() int {
	return len(p)
}

func (p PYS) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ByP struct {
	PYS
}

func (b ByP) Less(i, j int) bool {
	return b.PYS[i].P < b.PYS[j].P
}

type ByY struct {
	PYS
}

func (b ByY) Less(i, j int) bool {
	return b.PYS[i].Y < b.PYS[j].Y
}

type ByID struct {
	PYS
}

func (b ByID) Less(i, j int) bool {
	return b.PYS[i].Id < b.PYS[j].Id
}

var group = make(map[int][]PY)
var (
	n int
	m int
)

func main() {
	sc.Split(bufio.ScanWords)

	n, m = nextInt(), nextInt()
	for i := 0; i < m; i++ {
		p, y := nextInt(), nextInt()
		group[p] = append(group[p], PY{Id: i, P: p, Y: y, Rank: 0, Serial: ""})
	}

	// fmt.Println(group)

	ans := make([]PY, 0)

	for _, v := range group {
		sort.Sort(ByY{v})
		for i := 0; i < len(v); i++ {
			v[i].Rank = i + 1
			v[i].NewSerial()
			ans = append(ans, v[i])
		}
	}

	// fmt.Println(group)

	sort.Sort(ByID{ans})
	for _, a := range ans {
		fmt.Println(a.Serial)
	}
}
