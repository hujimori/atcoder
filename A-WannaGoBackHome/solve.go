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
	// sc.Split(bufio.ScanWords)

	var S string
	fmt.Scan(&S)
	m := map[string]int{"N": 0, "W": 0, "S": 0, "E": 0}

	for _, ss := range S {
		m[string(ss)]++
	}

	if (m["N"] == m["S"]) && (m["W"] == m["E"]) {
		fmt.Println("Yes")
	} else if (m["N"] > 0 && 0 < m["S"]) && (m["W"] == 0 && m["E"] == 0) {
		fmt.Println("Yes")
	} else if (m["N"] == 0 && m["S"] == 0) && (m["W"] > 0 && 0 < m["E"]) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
