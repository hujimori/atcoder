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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

type AB struct {
	A int
	B int
}
type ABS []AB

func (ab ABS) Len() int {
	return len(ab)
}

func (ab ABS) Swap(i, j int) {
	ab[i], ab[j] = ab[j], ab[i]
}

type ByA struct {
	ABS
}

func (b ByA) Less(i, j int) bool {
	return b.ABS[i].A < b.ABS[j].A
}

type ByB struct {
	ABS
}

func (b ByB) Less(i, j int) bool {
	return b.ABS[i].B < b.ABS[j].B
}

// func (b ByA) Less(i, j int) bool {
// if b[i].B < b[j].B {
// 	return true
// } else if b[i].B == b[i].B {
// 	if b[i].A < b[j].A {
// 		return true
// 	}
// }
// return false
// }

// type ByA []AB

// func (b ByA) Len() int {
// 	return len(b)
// }

// func (b ByA) Swap(i, j int) {
// 	b[i], b[j] = b[j], b[i]
// }

// func (b ByA) Less(i, j int) bool {
// 	// if b[i].B < b[j].B {
// 	// 	return true
// 	// } else if b[i].B == b[i].B {
// 	// 	if b[i].A < b[j].A {
// 	// 		return true
// 	// 	}
// 	// }
// 	// return false
// }

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	ab := make([]AB, n)
	for i := 0; i < n; i++ {
		ab[i] = AB{A: nextInt(), B: nextInt()}
	}

	// fmt.Println(ab)
	sort.Sort(ByA{ab})
	sort.Sort(ByB{ab})

	total := 0
	for i := 0; i < len(ab); i++ {
		total += ab[i].A

		if total <= ab[i].B {
			continue
		} else {
			fmt.Println("No")
			return
		}

	}

	fmt.Println("Yes")
}
