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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	slice := make([]int, n)

	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}

	length := 0
	start := 0
	for i := 0; i < len(slice)-1; i++ {
		// tmp := slice[start:len(slice)]
		fmt.Println(slice)

		for j := 1; j < len(slice); j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			if slice[i] < slice[j] && slice[j] < slice[j+1] {
				continue
			}

			if slice[i] > slice[j] && slice[j] < slice[j+1] {
				length++
				i = j
				break
			}

			if slice[i] > slice[j] && slice[j] > slice[j+1] {
				continue
			}

			if slice[i] < slice[j] && slice[j] > slice[j+1] {
				length++
				i = j
				break
			}

			if slice[j] == slice[j+1] {
				continue
			}

		}
		fmt.Println(length)

	}

	fmt.Println(length)

}
