package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	N := nextInt()

	s := make([][]string, 3)
	for i := 0; i < N; i++ {
		s[i] = strings.Split(nextLine(), "")
		sort.Strings(s[i])

	}

	fmt.Println(s)

	// s := make([]map[string]int, N)
	// for i := 0; i < len(s); i++ {
	// 	s[i] = make(map[string]int)
	// }

	// for i := 0; i < N; i++ {
	// 	str := strings.Split(nextLine(), "")
	// 	for _, st := range str {
	// 		if _, ok := s[i][st]; !ok {
	// 			s[i][st] = 1
	// 		} else {
	// 			s[i][st]++
	// 		}

	// 	}
	// }

	// cnt := 0
	// for i := 0; i < N-1; i++ {
	// 	for j := i + 1; j < N; j++ {
	// 		if reflect.DeepEqual(s[i], s[j]) {
	// 			cnt++
	// 		}
	// 	}
	// }

	// fmt.Println(cnt)

}
