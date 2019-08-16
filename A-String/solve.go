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

func main() {
	sc.Split(bufio.ScanWords)

	var X string
	fmt.Scan(&X)

	s := NewStack()

	for _, xx := range X {
		if xx == 'S' {
			s.Push(xx)
		}

		if xx == 'T' && (s.Peek() == 'T' || s.IsEmpty()) {
			s.Push(xx)
		}

		if xx == 'T' && s.Peek() == 'S' {
			s.Pop()
		}
	}

	ans := 0
	for !s.IsEmpty() {
		s.Pop()
		ans++
	}
	fmt.Println(ans)
}
