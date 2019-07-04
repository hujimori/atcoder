package main

import "fmt"

type Stack []int

func NewStack() Stack {
	s := make(Stack, 0)
	return s
}

func (s *Stack) push(i int) {
	*s = append(*s, i)
}

// 入力はすべて正とする
// 素数判定O(√n)
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	// 1の場合は例外
	return n != 1
}

// 約数の列挙O(√n)
func divisor(n int) Stack {
	s := NewStack()
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			s.push(i)
			if i != n/i {
				s.push(n / i)
			}
		}
	}
	return s
}

// 素因数分解
func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

func main() {

	if isPrime(11) {
		fmt.Printf("%d is prime\n", 11)
	}

	if !isPrime(12) {
		fmt.Printf("%d is not prime\n", 12)
	}

	res := divisor(12)
	fmt.Println(res)

	m := primeFactor(12)
	fmt.Println(m)
}
