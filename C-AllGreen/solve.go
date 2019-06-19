package main

import (
	"bufio"
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

var (
	d int
	g int
)

func main() {
	sc.Split(bufio.ScanWords)

	d = nextInt()
	g = nextInt()

	p := make([]int, d)
	c := make([]int, g)
	for i := 0; i < d; i++ {
		p[i] = nextInt()
		c[i] = nextInt()
	}

 for {

	for i := 0; i < d; i++{
	
		for j := 0; j < p[i]; j++ {
			sum += j * c[i]
		}


	}
}
	

}

func dfs(, score, count int) (int, int) {

	if i == d {
		return 0, score
	}


	dfs(i+1, score+)

}
