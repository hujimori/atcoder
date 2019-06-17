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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := nextInt()
	b := nextInt()

	if n == 2 {

	}

	// if n == 2 && a%2 == 1 && b%2 == 0 {
	// 	fmt.Println("Draw")
	// 	return
	// }

	if a%2 == 0 && b%2 == 0 {
		fmt.Println("Alice")
	} else if a%2 == 1 && b%2 == 0 {
		fmt.Println("Borys")
	} else if a%2 == 1 && b%2 == 1 {
		fmt.Println("Alice")
	} else if a%2 == 0 && b%2 == 1 {
		fmt.Println("Borys")
	}
}
