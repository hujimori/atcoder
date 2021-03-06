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

type Pair struct {
	p   int64
	num int64
}

func primeFactorize(n int64) *Stack {
	s := NewStack()
	for p := int64(2); p*p <= n; p++ {
		if n%p != 0 {
			continue
		}
		num := int64(0)
		for n%p == 0 {
			num++
			n /= p
		}
		s.Push(Pair{p: p, num: num})
	}
	if n != 1 {
		s.Push(Pair{n, 1})
	}
	return s
}

const MAX = 210000
const MOD = 1000000007

var fac, finv, inv [MAX]int64

func COMinit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := int64(2); i < MAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

func COM(n, k int64) int64 {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

var W, H int64

func main() {
	sc.Split(bufio.ScanWords)

	W, H = int64(nextInt()), int64(nextInt())
	COMinit()
	ans := COM(H+W-2, H-1)
	fmt.Println(ans)
}
