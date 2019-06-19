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

// var end int
// var slice []string

func main() {
	sc.Split(bufio.ScanWords)

	// slice = strings.Split(nextLine(), "")
	// end = len(slice)

	s := nextLine()

	// sum := 0
	// for i := 1; i < len(s); i++ {
	// 	s1 := s[0:i]
	// 	fmt.Printf("s1=%s ", s1)
	// 	a, _ := strconv.Atoi(s1)
	// 	sum += a
	// 	for j := i + 1; j < len(s); j++ {
	// 		s2 := s[i:j]
	// 		fmt.Printf("s2=%s ", s2)
	// 		b, _ := strconv.Atoi(s2)
	// 		sum += b
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(sum)

	// fmt.Println(s)
	// fmt.Println(s[0])
	// fmt.Println(s[0:1])
	// fmt.Println(s[0:2])
	// fmt.Println(s[1:len(s)])
	ans := dfs(0, s)
	fmt.Println(ans)
}

func dfs(sum int, s string) int {
	result := sum
	for i := 1; i < len(s); i++ {
		a, _ := strconv.Atoi(s[0:i])
		result += dfs(sum+a, s[i:len(s)])
	}
	a, _ := strconv.Atoi(s)
	return result + a
}

// func dfs(sum int, s string) int {

// 	if len(s) == 0 {
// 		return sum
// 	}

// if i == len(s)+1 {
// 	return sum
// }

// i番目の後ろに+を挿入する場合
// var sum1 int
// s1, _ := strconv.Atoi(s[0:1])

// fmt.Printf("sum+s1=%d sum=%d s1=%d, s[i:len(s)]=%s\n", sum+s1, sum, s1, s[1:len(s)])
// sum1 += dfs(sum+s1, s[1:len(s)])

// i番目の後ろに+を挿入しない場合
// var sum2 int
// // s2, _ := strconv.Atoi(s[0:len(s)])
// fmt.Printf("sum+s2=%d  sum=%d s2=%d, s[0:len(s)]=%s\n", sum+0, sum, 0, s[0:len(s)])
// sum2 += dfs(sum, s[0:len(s)])

// return sum2
// }
