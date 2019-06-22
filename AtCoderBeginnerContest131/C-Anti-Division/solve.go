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

	a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
	origc := c
	origd := d
	x := c * d
	// if c < d {
	// 	tmp := c
	// 	c = d
	// 	d = tmp
	// }

	r := c % d
	for r != 0 {
		c = d
		d = r
		r = c % d
	}

	// fmt.Println(x / d)

	total := 0
	ans := b - a + 1
	// fmt.Printf("ans=%d\n", ans)
	for end := b; end >= a; end-- {
		if end%origc == 0 {
			total += (end/origc - (a-1)/origc)

			break
		}
	}

	for end := b; end >= a; end-- {
		if end%origd == 0 {
			total += (end/origd - (a-1)/origd)
			break
		}
	}

	gcd := x / d

	for end := b; end >= a; end-- {
		// fmt.Println(end)
		if end%gcd == 0 {
			total -= end/gcd - (a-1)/gcd
			break
		} else {
			end -= end % gcd

		}
	}

	ans -= total
	fmt.Println(ans)

}
