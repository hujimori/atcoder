package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	sc.Split(bufio.ScanWords)

	sx, sy, tx, ty := nextInt(), nextInt(), nextInt(), nextInt()

	dx := tx - sx
	dy := ty - sy

	ans := ""

	// Path1
	ans += strings.Repeat("U", dy) + strings.Repeat("R", dx)

	// Path2
	ans += strings.Repeat("D", dy) + strings.Repeat("L", dx)

	// Path3
	ans += "L" + strings.Repeat("U", dy+1) + strings.Repeat("R", dx+1) + "D"

	// Path4
	ans += "R" + strings.Repeat("D", dy+1) + strings.Repeat("L", dx+1) + "U"

	fmt.Println(ans)
}
