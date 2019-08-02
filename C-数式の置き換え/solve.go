package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type StackData struct {
	data interface{}
	next *StackData
}

type Stack struct {
	sp  *StackData
	top int
}

func NewStack() *Stack {
	return &Stack{top: 0}
}

func (s *Stack) Push(data interface{}) {
	s.sp = &StackData{data: data, next: s.sp}
	s.top++
}

func (s *Stack) Pop() interface{} {
	if s.top > 0 {
		item := s.sp.data
		s.sp = s.sp.next
		s.top--
		return item
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.top > 0 {
		return s.sp.data
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return s.top
}

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

	var S string
	fmt.Scan(&S)

	ans := 0

	flag := true
	for i := 0; i < len(S); i++ {
		if string(S[i]) == "+" {
			if flag {
				ans++
			}
			flag = true
		} else if string(S[i]) == "0" {
			flag = false
		}
	}

	if flag {
		ans++
	}

	fmt.Println(ans)

}
