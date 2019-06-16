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
	count := 0

	for i := 0; i < n-1; i++ {
		sum := slice[i]
		for j := i + 1; j < n; j++ {
			fmt.Printf("%d + %d = %d\n", sum, slice[j], sum+slice[i])
			sum += slice[j]
			fmt.Printf("sum = %d\n", sum)
			if sum >= k {
				count += n - j
				fmt.Printf("count %d = %d - %d\n", count, n, j)

				break

			}
		}
	}

	// if slice[len(slice)-1] >= k {
	// 	count++
	// }

	fmt.Println(count)
}
