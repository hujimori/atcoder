package main

import "fmt"

var isPrime []bool
var isPrimeSmall []bool

// [a, b)の整数に対して篩を行う isPrime[i-a]  trye <=> iが素数
func segmentSieve(a, b uint64) {
	isPrime = make([]bool, b-a)
	isPrimeSmall = make([]bool, b)

	for i := uint64(0); i*i < b; i++ {
		isPrimeSmall[i] = true
	}
	for i := uint64(0); i < b-a; i++ {
		isPrime[i] = true
	}

	for i := uint64(2); i*i < b; i++ {
		if isPrimeSmall[i] {
			for j := 2 * i; j*j < b; j += i {
				isPrimeSmall[j] = false
			}
			for j := max(2, (a+i-1)/i) * i; j < b; j += i {
				isPrime[j-a] = false
			}
		}
	}

}

func main() {
	var (
		a uint64
		b uint64
	)
	fmt.Scan(&a)
	fmt.Scan(&b)

	segmentSieve(a, b)

	s := make([]uint64, 0)
	for i, v := range isPrime {
		if v {
			s = append(s, uint64(i)+a)
		}
	}

	fmt.Println(s)
	fmt.Println(len(s))

}

func max(a, b uint64) uint64 {
	if a < b {
		return b
	}
	return a
}
