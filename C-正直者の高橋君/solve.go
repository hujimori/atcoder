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

type QueueData struct {
	data interface{}
	prev *QueueData
}

type Queue struct {
	front *QueueData
	last  *QueueData
	top   int
}

func NewQueue() *Queue {
	return &Queue{top: 0}
}

func (q *Queue) Enqueue(data interface{}) {
	qd := &QueueData{data: data, prev: nil}
	if q.top == 0 {
		q.front = qd
		q.last = q.front
		q.top++
		return
	}

	q.last.prev = qd
	q.last = qd
	q.top++
}

func (q *Queue) Dequeue() interface{} {
	if q.top > 0 {
		d := q.front.data
		q.front = q.front.prev
		q.top--

		return d
	}
	return nil
}

func (q *Queue) Front() interface{} {
	return q.front.data
}

func (q *Queue) Last() interface{} {
	return q.last.data
}

func (q *Queue) IsEmpty() bool {
	if q.top == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.top
}

var G [][]int
var N int
var M int
var a, b int
var d [310]int
var memo [310]int

const MOD = 1e9 + 7
const INF = 1e9

func main() {
	sc.Split(bufio.ScanWords)

	N = nextInt()
	for i := 0; i < N; i++ {
		d[i] = INF
	}

	a, b = nextInt()-1, nextInt()-1
	M = nextInt()

	G = make([][]int, N)
	for i := 0; i < M; i++ {
		u, v := nextInt()-1, nextInt()-1

		G[u] = append(G[u], v)
		G[v] = append(G[v], u)

	}

	q := NewQueue()
	q.Enqueue(a)
	d[a] = 0
	memo[a] = 1

	for !q.IsEmpty() {
		u := q.Dequeue().(int)
		if u == b {
			break
		}

		for _, v := range G[u] {
			if d[v] == INF {
				d[v] = d[u] + 1
				memo[v] += memo[u]
				memo[v] %= MOD
				q.Enqueue(v)
			} else if d[v] == d[u]+1 {
				memo[v] += memo[u]
				memo[v] %= MOD
			}
		}
	}
	fmt.Println(memo[b] % MOD)
}
