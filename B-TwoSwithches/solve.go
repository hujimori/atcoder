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

var grid [][]string

type alice struct {
	A int
	B int
}

type bob struct {
	C int
	D int
}

func main() {
	sc.Split(bufio.ScanWords)

	alice := alice{nextInt(), nextInt()}
	bob := bob{nextInt(), nextInt()}

	max := math.Max(float64(alice.A), float64(bob.C))
	min := math.Min(float64(alice.B), float64(bob.D))

	if max <= min {
		fmt.Println(int(min - max))
	} else {
		fmt.Println(0)
	}

}

// func main() {
// 	sc.Split(bufio.ScanWords)

// 	alice := alice{nextInt(), nextInt()}
// 	bob := bob{nextInt(), nextInt()}

// 	dt := 0

// 	if alice.A <= bob.C && bob.C <= alice.B && alice.A <= bob.D && bob.D < alice.B {
// 		dt = bob.D - bob.C
// 	} else if bob.C <= alice.A && alice.A <= bob.D && bob.C <= alice.B && alice.B <= bob.D {
// 		dt = alice.B - alice.A
// 	} else if alice.A <= bob.C && bob.C <= alice.B && alice.B <= bob.D {
// 		dt = alice.B - bob.C
// 	} else if alice.A <= bob.D && bob.D <= alice.B && bob.C <= alice.A {
// 		dt = bob.D - alice.A
// 	}

// 	fmt.Println(dt)

// }
