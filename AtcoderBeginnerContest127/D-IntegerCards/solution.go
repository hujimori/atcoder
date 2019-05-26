package main

import "fmt"

type BC struct {
	b int
	c int
}

func main() {

	var n int
	fmt.Scan(&n)
	inputA := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&inputA[i])
	}

	var m int
	fmt.Scan(&m)
	inputBC := make([]BC, m, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&inputBC[i])
	}

	max := 0

	for i := 0; i < m; i++ {
		value := 0
		for j := 0; j < inputBC[i].b; j++ {
			value += inputBC[i].c
		}
		for k := 0; k < n; k++ {
			if max < value+inputA[k] {
				max = value + inputA[k]
			}
		}
	}

	fmt.Println(max)

}
