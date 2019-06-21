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

type AB struct {
	A int
	B int
}

type ABS []AB

func (abs ABS) Len() int {
	return len(abs)
}

func (abs ABS) Swap(i, j int) {
	abs[i], abs[j] = abs[j], abs[i]
}

func (abs ABS) Less(i, j int) bool {
	return abs[i].B-abs[i].A < abs[j].B-abs[j].A
}

// type ByA struct {
// 	ABS
// }

// func (b ByA) Less(i, j int) bool {
// 	return b.ABS[i].A < b.ABS[j].A
// }

// type ByB struct {
// 	ABS
// }

// func (b ByB) Less(i, j int) bool {
// 	return b.ABS[i].B < b.ABS[j].B
// }

func main() {
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	slice := make([]AB, m)
	for i := 0; i < m; i++ {
		slice[i] = AB{A: nextInt(), B: nextInt()}
	}

	// sort.Sort(ByA{slice})
	// // sort.Sort(ByB{slice})
	// fmt.Println(bridge)

	// sort.Sort(ABS(slice))

	// startA := slice[0].A
	// startB := slice[0].B

	// 切り落とす橋の本数
	cutBridge := 1

	for i := 0; i < m-1; i++ {
		if slice[i].A <= slice[i+1].B && slice[i].B >= slice[i+1].A {
			continue
		} else {
			cutBridge++
		}
	}

	fmt.Println(cutBridge)
}
