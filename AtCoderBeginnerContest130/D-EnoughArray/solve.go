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
	k := nextInt()

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	answer := 0
	sum := 0

	r := 0

	for l := 0; l < n; l++ {
		for sum < k {
			if r == n {
				break
			} else {
				sum += slice[r]
				r++
			}
		}

		if sum < k {
			break
		}

		answer += n - r + 1
		sum -= slice[l]

	}

	fmt.Println(answer)
	// var sum int
	// for i := 0; i < n; i++ {
	// 	slice[i] = nextInt()
	// 	sum += slice[i]
	// }
	// count := 0

	// sumA := 0

	// for i := 0; i < n; i++ {

	// 	sumA += slice[i]

	// 	if sum-sumA < k {
	// 		count += i

	// 		fmt.Printf("count %d\n", count)
	// 		break

	// 	}

	// }

	// sumB := 0
	// for i := 0; i < n; i++ {
	// fmt.Printf("%d + %d = %d\n", sumB, slice[i], sumB+slice[i])

	// 	sumB += slice[i]

	// 	if sumB >= k {
	// 		fmt.Printf("count %d = %d - %d\n", n-i, n, i)

	// 		count += n - i
	// 		break
	// 	}

	// }

	// for i := 0; i < n-1; i++ {
	// 	sum := 0
	// 	for j := i; j < n; j++ {
	// 		// fmt.Printf("%d + %d = %d\n", sum, slice[j], sum+slice[j])
	// 		sum += slice[j]
	// 		// fmt.Printf("sum = %d\n", sum)
	// 		if sum >= k {
	// 			count += n - j
	// 			// fmt.Printf("count %d = %d - %d\n", count, n, j)
	// 			break
	// 		}
	// 	}
	// }

	// if slice[len(slice)-1] >= k {
	// 	count++
	// }

	// fmt.Println(count)
}
