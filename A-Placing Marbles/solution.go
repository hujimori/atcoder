package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	count := 0

	for _, s := range str {
		if s == '1' {
			count++
		}
	}

	fmt.Println(count)
}
