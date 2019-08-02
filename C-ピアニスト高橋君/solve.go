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
	var s string
	fmt.Scan(&s)
	kenban := "WBWBWWBWBWBWWBWBWWBWBWBWWBWBWWBWBWBWWBWBWWBWBWBWWBWBWWBWBWBWWBWBWWBWBWBWWBWBWWBWBWWBWBWBWWBWBW"

	white := 0
	for i := 0; i < len(kenban); i++ {
		if kenban[i:len(s)+i] == s {
			break
		}

		if string(kenban[i]) == "W" {
			white++
		}
	}

	pos := []string{"Do", "Re", "Mi", "Fa", "So", "La", "Si"}

	fmt.Println(pos[white])
}
