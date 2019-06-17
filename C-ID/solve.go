package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type PY struct {
	Id int
	P  int
	Y  int
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

type Serial struct {
	Id     int
	Number int
}

type Serials []Serial

func (s Serials) Len() int {
	return len(s)
}

func (s Serials) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Serials) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	m := nextInt()

	slice := make([]PY, m)
	for i := 0; i < m; i++ {
		slice[i].Id = i
		slice[i].P = nextInt()
		slice[i].Y = nextInt()
	}

	fmt.Println(n)

	fmt.Println(slice)

	sort.Sort(ByP{slice})
	sort.Sort(ByY{slice})

	fmt.Println(slice)

	serials := make([]Serial, m)
	for i, s := range slice {
		upper := strconv.Itoa(s.P)
		for len(upper) == 6 {
			upper = "0" + upper
		}

		lower := strconv.Itoa(s.Y)
		for len(lower) == 6 {
			lower = "0" + lower
		}

		serials[i].Id = s.Id
		serials[i].Number, _ = strconv.Atoi(upper + lower)

	}

	// sort.Sort(serials)
	for _, s := range serials {
		fmt.Println(s.Number)
	}
}
