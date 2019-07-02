package main

import (
	"bufio"
	"container/heap"
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

type Item struct {
	L     int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].L < pq[j].L
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Fence Rapair
var priorityQueue PriorityQueue
var n int
var L []int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	L = make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = nextInt()
	}

	priorityQueue = make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		priorityQueue[i] = &Item{L: L[i]}
		priorityQueue[i].Index = i
	}

	heap.Init(&priorityQueue)

	solve()
}

func solve() {
	ans := 0
	for priorityQueue.Len() > 1 {
		// 一番短い板, 次に短い板を取り出す
		l1, l2 := heap.Pop(&priorityQueue).(*Item).L, heap.Pop(&priorityQueue).(*Item).L
		ans += l1 + l2

		heap.Push(&priorityQueue, &Item{L: ans})
	}

	println(ans)
}
