package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Gas   int
	Index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Gas > pq[j].Gas
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

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

var priorityQueue PriorityQueue
var (
	N int
	L int
	P int
)
var A []int
var B []int

func main() {
	sc.Split(bufio.ScanWords)
	N, L, P = nextInt(), nextInt(), nextInt()
	A = make([]int, N+10)
	B = make([]int, N+10)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
		B[i] = nextInt()
	}

	priorityQueue = make(PriorityQueue, N)
	for i := 0; i < N; i++ {
		priorityQueue[i] = &Item{}
		priorityQueue[i].Index = i
	}

	heap.Init(&priorityQueue)

	solve()

}

func solve() {
	A[N] = L
	B[N] = 0
	N++

	// ans:補填回数, pos:今いる場所, tan:タンクのガソリン量
	ans, pos, tank := 0, 0, P

	for i := 0; i < N; i++ {
		// 次に進む距離
		d := A[i] - pos

		// 十分な量になるまでガソリンを補給
		for tank-d < 0 {
			if priorityQueue.Len() == 0 {
				fmt.Println("-1")
				return
			}

			tank += heap.Pop(&priorityQueue).(*Item).Gas
			ans++
		}

		tank -= d
		pos = A[i]
		heap.Push(&priorityQueue, &Item{Gas: B[i]})

	}

	fmt.Println(ans)
}
