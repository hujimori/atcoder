package main

import "fmt"

func main() {

	// bitSearch()
	// solve()
	// solve2()
	solve3()
}

func bitSearch() {

	n := uint(5)

	for bit := uint(0); bit < (1 << n); bit++ {

		s := []uint{}

		for i := uint(0); i < n; i++ {
			if bit&(1<<i) == 1<<i {
				s = append(s, i)
			}

		}

		fmt.Printf("%d: {", bit)
		for i := 0; i < len(s); i++ {
			fmt.Printf(" %d ", s[i])
		}
		fmt.Printf("}\n")
	}
}

func solve() {
	s := []int{7, 2, 9}
	n := uint(3)
	k := 11
	exist := false
	var slice []int
	for bit := uint(0); bit < (1 << n); bit++ {
		sum := 0
		slice = []int{}
		for i := uint(0); i < n; i++ {
			if bit&(1<<i) == 1<<i {
				sum += s[i]
				slice = append(slice, s[i])
			}
		}

		if sum == k {
			exist = true
		}
	}

	fmt.Println(slice)
	if exist {

		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func solve2() {
	n := uint(10)
	a := (1 << 2) | (1 << 3) | (1 << 5) | (1 << 7)

	for bit := a; ; bit = (bit - 1) & a {

		s := []uint{}
		for i := uint(0); i < n; i++ {
			if bit&(1<<i) == 1<<i {
				s = append(s, i)
			}
		}

		fmt.Printf("%d: {", bit)
		for i := 0; i < len(s); i++ {
			fmt.Printf("%d ", s[i])
		}
		fmt.Printf("}\n")

		if bit == 0 {
			break
		}
	}
}

func nextCombination(sub int) int {
	x := sub & -sub
	y := sub + x
	return (((sub & ^y) / x) >> 1) | y
}

func solve3() {
	n := 5
	k := 3

	bit := (1<<uint(k) - 1)
	for ; bit < (1 << uint(n)); bit = nextCombination(bit) {

		slice := []int{}
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) == 1<<uint(i) {
				slice = append(slice, i)
			}
		}

		fmt.Printf("%d: {", bit)
		for i := 0; i < len(slice); i++ {
			fmt.Printf("%d ", slice[i])
		}
		fmt.Print("}\n")

	}
}
