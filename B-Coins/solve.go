package main

import "fmt"

func main() {

	var A int
	fmt.Scan(&A)

	var B int
	fmt.Scan(&B)

	var C int
	fmt.Scan(&C)

	var X int
	fmt.Scan(&X)

	count := 0

	for i := 0; i < A+1; i++ {
		for j := 0; j < B+1; j++ {
			for k := 0; k < C+1; k++ {
				if 500*i+100*j+50*k == X {
					count++
				}
			}
		}
	}

	fmt.Println(count)

}
