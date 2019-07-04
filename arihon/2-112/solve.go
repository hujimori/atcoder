package main

import "fmt"

// i番目の素数
var prime []int

// isPrime[i]がtrueならiは素数
var isPrime []bool

// n以下の素数を返す
func sieve(n int) int {

	prime = make([]int, n)
	isPrime = make([]bool, n+1)

	p := 0
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime[p] = i
			p++
			for j := 2 * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	return p
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%d以下の素数の個数 = %d\n", n, sieve(n))
}
