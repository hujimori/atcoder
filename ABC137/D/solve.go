package main

import (
	"bufio"
	"container/heap"
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

type PriorityQueue []int

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i] > pq[j] }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := nextInt(), nextInt()
	ts := make(map[int][]int)

	for i := 0; i < N; i++ {
		a, b := nextInt(), nextInt()
		ts[a] = append(ts[a], b)
	}

	pq := make(PriorityQueue, 0, N)
	heap.Init(&pq)
	ans := 0
	for i := 1; i <= M; i++ {
		if _, ok := ts[i]; ok {
			for _, v := range ts[i] {
				heap.Push(&pq, v)
			}
		}
		if len(pq) > 0 {
			ans += heap.Pop(&pq).(int)
		}
	}
	fmt.Println(ans)
}
