package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	m := map[string]int{}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		m[s[:1]]++
	}

	march := []string{"M", "A", "R", "C", "H"}

	ans := 0
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			for k := j + 1; k < 5; k++ {
				ans += m[march[i]] * m[march[j]] * m[march[k]]
			}
		}
	}

	fmt.Println(ans)
}
